package main

import (
	"discord/service"

	"github.com/bwmarrin/discordgo"
)

// commandHandlers 는 [string]func(s,i) map 이므로, commands 에서 참조로 ApplicationComand 를 수정했기에 Name 을 조회하며 함수가 있으면 ok가 되어 핸들러 등록이 되는듯.
func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) { //Interaction 이 생성되었을 떄의 handler 을 정의한다.

	if i.Type == discordgo.InteractionApplicationCommand { //슬래시 커맨드를 입력받았을 때
		if handlerCommand, ok := commandHandlers[i.ApplicationCommandData().Name]; //맵 조회 패턴. handler 는 함수가 들어갈것이고 ok 에는 맵에 값이 있는지를 표현하는 불리언값임
		ok {
			handlerCommand(s, i)
		}
	}

	if i.Type == discordgo.InteractionMessageComponent {

		switch i.MessageComponentData().CustomID {

		//좌석 관련 버튼
		case "sitdown_btn":
			service.TakeaSeat(s, i)
		case "standup_btn":
			service.Standup(s, i)
		case "update_btn":
			service.UpdateRoomState(s, i)

		//문 관련 버튼
		case "opendoor_btn":
			service.OpentheDoor(s, i)
		case "closedoor_btn":
			service.ClosetheDoor(s, i)
		}


	}
}
