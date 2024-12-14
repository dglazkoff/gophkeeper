package auth

import (
	"context"
	"fmt"
	"gophkeeper/internal/logger"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

const TokenExp = time.Hour * 3
const SecretKey = "supersecretkey"

func authMethods() map[string]bool {
	const storageServicePath = "/models.Storage/"

	return map[string]bool{
		storageServicePath + "SavePassword":   true,
		storageServicePath + "GetPassword":    true,
		storageServicePath + "DeletePassword": true,

		storageServicePath + "SaveText":   true,
		storageServicePath + "GetText":    true,
		storageServicePath + "DeleteText": true,

		storageServicePath + "SaveBinary":   true,
		storageServicePath + "GetBinary":    true,
		storageServicePath + "DeleteBinary": true,

		storageServicePath + "SaveBankCard":   true,
		storageServicePath + "GetBankCard":    true,
		storageServicePath + "DeleteBankCard": true,
	}
}

type InterceptorServer struct {
}

func NewInterceptorServer() *InterceptorServer {
	return &InterceptorServer{}
}

func BuildJWTString(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func extractUserIDFromToken(token string) (string, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		logger.Log.Error("Error while parse token: ", err)
		return "", err
	}

	return claims.UserID, nil
}

func GetUserIdFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		logger.Log.Error("userId not found in context")
		return "", fmt.Errorf("userId not found in context")
	}
	return userID, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	logger.Log.Debug("Valid token")
	return nil
}

func (i *InterceptorServer) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if isAuth := authMethods()[info.FullMethod]; !isAuth {
			return handler(ctx, req)
		}

		logger.Log.Debug("Unary interceptor: ", info.FullMethod)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := md["authorization"]
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		accessToken := values[0]
		err := verifyToken(accessToken)

		if err != nil {
			logger.Log.Error("Error while verify token: ", err)
			return nil, status.Errorf(codes.Unauthenticated, "access token is invalid")
		}

		userID, err := extractUserIDFromToken(accessToken)

		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "user is not defined")
		}

		ctx = context.WithValue(ctx, "userID", userID)
		return handler(ctx, req)
	}
}
