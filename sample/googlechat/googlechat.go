package main

import (
	"context"
	"fmt"

	// https://github.com/googleapis/google-api-go-client
	// "cloud.google.com/go/storage"

	chat "google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

// errors
const (
	ERR_NEW_SERVICE_F      = "Service Account 정보를 통해 서비스 생성 중 에러 : %s"
	ERR_CALL_SPACE_F       = "지정한 Space로 메시지 전송 중 에러 : %s"
	ERR_RESPONSE_MARSHAL_F = "응답 데이터를 json.marshal 중 에러 : %s"
)

// global vars
const (
	SCOPE    = "https://www.googleapis.com/auth/chat.bot"
	CHAT_MSG = "채팅 봇이에요, 반가워요"
	SPACE    = "spaces/__________" // google chat의 url 중 spaces/****/ 에 해당하는 부분을 복사
	SC_PATH  = "/____/____/_/____/sc.json"
	PROJ_ID  = "__PROJECT_ID__"
)

// sercie
var Byte = []byte(`
{
    "type": "service_account",
    "project_id": "__PROJECT_ID__",
    "private_key_id": "ad3f..............e9",
    "private_key": "-----BEGIN PRIVATE KEY-----\n........................\n-----END PRIVATE KEY-----\n",
    "client_email": "__SERVIECE_ACCOUNT_EMAIL__",
    "client_id": "__CLIENT_ID__",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/__SERVIECE_ACCOUNT_EMAIL__"
  }
  `) // client_x509_cert_url에서 __SERVIECE_ACCOUNT_EMAIL__의 값은 url encoding 해야한다.

// [START auth_cloud_explicit]
func main() {
	explicitChat()
}

// serviceaccount를 사용해 chatbot service에 접근해보자
func explicitChat() {
	var (
		ctx context.Context = context.Background()
		// token                *oauth2.Token                  = &oauth2.Token{}
		sErr                 error                          = nil
		chatService          *chat.Service                  = nil
		sSpaceMessageService *chat.SpacesMessagesService    = nil
		sMessage             *chat.Message                  = &chat.Message{Text: CHAT_MSG}
		sCall                *chat.SpacesMessagesCreateCall = nil
		sByte                []byte                         = nil
	)

	chatService, sErr = chat.NewService(ctx, option.WithCredentialsJSON(Byte), option.WithScopes(SCOPE))
	// chatService, sErr = chat.NewService(ctx, option.WithCredentialsFile(SC_PATH), option.WithScopes(Scope))
	if sErr != nil {
		fmt.Printf(ERR_NEW_SERVICE_F, sErr.Error())
		return
	}

	sSpaceMessageService = chat.NewSpacesMessagesService(chatService)
	sCall = sSpaceMessageService.Create(SPACE, sMessage)

	sMessage, sErr = sCall.Do()
	if sErr != nil {
		fmt.Printf(ERR_CALL_SPACE_F, sErr.Error())
		return
	}

	sByte, sErr = sMessage.MarshalJSON()
	if sErr != nil {
		fmt.Printf(ERR_RESPONSE_MARSHAL_F, sErr.Error())
		return
	}
	fmt.Println(string(sByte))
	return
}
