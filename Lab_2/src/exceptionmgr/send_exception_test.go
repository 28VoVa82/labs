package exceptionmgr

import (
	"Lab_2/src/cfg"
	"Lab_2/src/exchange"
	"Lab_2/src/exchange/mock_exchange"
	"github.com/golang/mock/gomock"
	"net"
	"testing"
)

func TestExceptionSender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cliCfg := cfg.ClientConfig{
		CriticalErrorsList: []string{"LessThanZero"},
		SrvCred:            cfg.ServerCredentials{},
	}

	var c net.Conn
	m := mock_exchange.NewMockNetworkCommunication(ctrl)
	m.EXPECT().DialServer(gomock.Eq(&cliCfg)).Return(c).AnyTimes()
	m.EXPECT().CloseConnection(gomock.Eq(c)).Return().AnyTimes()

	t.Run("Exception sending okay", func(t *testing.T) {
		request := exchange.SerializeRequest(exchange.ClientRequest{
			ErrType:    "LessThanZero",
			IsCritical: true,
		})

		respond := exchange.SerializeRespond(exchange.ServerRespond{IsSuccess: true})

		m.EXPECT().SendData(gomock.Eq(request), gomock.Eq(c)).Return(true).AnyTimes()
		m.EXPECT().ReceiveData(gomock.Eq(c)).Return(respond).AnyTimes()
		sender := exchange.NewNetWorker(&cliCfg, m)
		errCnt := NewErrorCounter(sender, cliCfg.CriticalErrorsList)
		errCnt.SendException(&LessThanZero{})

		if errCnt.BadSending != 0 {
			t.Errorf("Expected `0` value. Actual value %d", errCnt.BadSending)
		}
	})

	t.Run("Exception sending failed", func(t *testing.T) {
		request := exchange.SerializeRequest(exchange.ClientRequest{
			ErrType:    "LessThanZero",
			IsCritical: true,
		})

		m.EXPECT().SendData(gomock.Eq(request), gomock.Eq(c)).Return(false).AnyTimes()
		sender := exchange.NewNetWorker(&cliCfg, m)
		errCnt := NewErrorCounter(sender, cliCfg.CriticalErrorsList)
		errCnt.SendException(&LessThanZero{})

		if errCnt.BadSending == 1 {
			t.Errorf("Expected `1` value. Actual value %d", errCnt.BadSending)
		}
	})
}

