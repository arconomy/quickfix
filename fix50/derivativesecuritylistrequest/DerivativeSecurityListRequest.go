//Package derivativesecuritylistrequest msg type = z.
package derivativesecuritylistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a DerivativeSecurityListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"z"`
	fixt11.Header
	//SecurityReqID is a required field for DerivativeSecurityListRequest.
	SecurityReqID string `fix:"320"`
	//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
	SecurityListRequestType int `fix:"559"`
	//UnderlyingInstrument is a non-required component for DerivativeSecurityListRequest.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
	SecuritySubType *string `fix:"762"`
	//Currency is a non-required field for DerivativeSecurityListRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for DerivativeSecurityListRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for DerivativeSecurityListRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
	SubscriptionRequestType *string `fix:"263"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized DerivativeSecurityListRequest instance
func New(securityreqid string, securitylistrequesttype int) *Message {
	var m Message
	m.SetSecurityReqID(securityreqid)
	m.SetSecurityListRequestType(securitylistrequesttype)
	return &m
}

func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = v }
func (m *Message) SetSecurityListRequestType(v int) { m.SecurityListRequestType = v }
func (m *Message) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *Message) SetSecuritySubType(v string)         { m.SecuritySubType = &v }
func (m *Message) SetCurrency(v string)                { m.Currency = &v }
func (m *Message) SetText(v string)                    { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)             { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)             { m.EncodedText = &v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "z", r
}
