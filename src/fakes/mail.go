package fakes

import "sync"

type MailIFace struct {
	SendEmailAuthCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Email string
			Code  string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
}

func (f *MailIFace) SendEmailAuth(param1 string, param2 string) error {
	f.SendEmailAuthCall.mutex.Lock()
	defer f.SendEmailAuthCall.mutex.Unlock()
	f.SendEmailAuthCall.CallCount++
	f.SendEmailAuthCall.Receives.Email = param1
	f.SendEmailAuthCall.Receives.Code = param2
	if f.SendEmailAuthCall.Stub != nil {
		return f.SendEmailAuthCall.Stub(param1, param2)
	}
	return f.SendEmailAuthCall.Returns.Error
}
