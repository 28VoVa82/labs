package exceptionmgr

import (
	"Lab_2/src/exchange"
	"log"
)

func (errCnt *ErrorCounter) SendException(err error) bool {
	if err == nil {
		return false
	}

	errMgr, isCastOk := err.(ErrMgr)
	if !isCastOk {
		return false
	}

	req := exchange.SerializeRequest(exchange.ClientRequest{
		ErrType:    errMgr.GetErrorType(),
		IsCritical: errCnt.isCritical(errMgr),
	})
	recvData := errCnt.Sender.SendRecvData(req)
	if recvData == nil {
		errCnt.BadSending++
		log.Print("Exception sending failed")
		return false
	}

	resp, isParseOk := exchange.ParseRespond(recvData)
	if !isParseOk {
		errCnt.BadSending++
		return false
	}

	if !resp.IsSuccess {
		errCnt.BadSending++
	}
	return true
}
