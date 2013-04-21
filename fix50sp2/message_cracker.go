package fix50sp2

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header.StringValue(tag.MsgType); msgType {
	case "6":
		return router.OnFIX50SP2IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX50SP2Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX50SP2ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50SP2OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "B":
		return router.OnFIX50SP2News(News{msg}, sessionID)
	case "C":
		return router.OnFIX50SP2Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX50SP2NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX50SP2NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX50SP2OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX50SP2OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX50SP2OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX50SP2AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "K":
		return router.OnFIX50SP2ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX50SP2ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX50SP2ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX50SP2ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX50SP2AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50SP2DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX50SP2QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX50SP2Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX50SP2SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX50SP2MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX50SP2MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX50SP2MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX50SP2MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX50SP2QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX50SP2QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50SP2MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX50SP2SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX50SP2SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX50SP2SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX50SP2SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX50SP2TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX50SP2TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX50SP2MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX50SP2BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX50SP2BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50SP2BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX50SP2ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX50SP2RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX50SP2RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX50SP2OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX50SP2OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX50SP2NewOrderCross(NewOrderCross{msg}, sessionID)
	case "t":
		return router.OnFIX50SP2CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "u":
		return router.OnFIX50SP2CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX50SP2SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX50SP2SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX50SP2SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX50SP2SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX50SP2DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX50SP2DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX50SP2NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX50SP2MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AD":
		return router.OnFIX50SP2TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX50SP2TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX50SP2OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX50SP2QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX50SP2RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX50SP2QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "AJ":
		return router.OnFIX50SP2QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AK":
		return router.OnFIX50SP2Confirmation(Confirmation{msg}, sessionID)
	case "AL":
		return router.OnFIX50SP2PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX50SP2PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "AN":
		return router.OnFIX50SP2RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX50SP2RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AP":
		return router.OnFIX50SP2PositionReport(PositionReport{msg}, sessionID)
	case "AQ":
		return router.OnFIX50SP2TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AR":
		return router.OnFIX50SP2TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50SP2AllocationReport(AllocationReport{msg}, sessionID)
	case "AT":
		return router.OnFIX50SP2AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AU":
		return router.OnFIX50SP2ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AV":
		return router.OnFIX50SP2SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX50SP2AssignmentReport(AssignmentReport{msg}, sessionID)
	case "AX":
		return router.OnFIX50SP2CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX50SP2CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX50SP2CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BA":
		return router.OnFIX50SP2CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX50SP2CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BC":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BD":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX50SP2UserRequest(UserRequest{msg}, sessionID)
	case "BF":
		return router.OnFIX50SP2UserResponse(UserResponse{msg}, sessionID)
	case "BG":
		return router.OnFIX50SP2CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BH":
		return router.OnFIX50SP2ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP2ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP2SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP2SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP2AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP2AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP2ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP2TradingSessionList(TradingSessionList{msg}, sessionID)
	case "BI":
		return router.OnFIX50SP2TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "BQ":
		return router.OnFIX50SP2SettlementObligationReport(SettlementObligationReport{msg}, sessionID)
	case "BR":
		return router.OnFIX50SP2DerivativeSecurityListUpdateReport(DerivativeSecurityListUpdateReport{msg}, sessionID)
	case "BS":
		return router.OnFIX50SP2TradingSessionListUpdateReport(TradingSessionListUpdateReport{msg}, sessionID)
	case "BT":
		return router.OnFIX50SP2MarketDefinitionRequest(MarketDefinitionRequest{msg}, sessionID)
	case "BU":
		return router.OnFIX50SP2MarketDefinition(MarketDefinition{msg}, sessionID)
	case "BV":
		return router.OnFIX50SP2MarketDefinitionUpdateReport(MarketDefinitionUpdateReport{msg}, sessionID)
	case "BW":
		return router.OnFIX50SP2ApplicationMessageRequest(ApplicationMessageRequest{msg}, sessionID)
	case "BX":
		return router.OnFIX50SP2ApplicationMessageRequestAck(ApplicationMessageRequestAck{msg}, sessionID)
	case "BY":
		return router.OnFIX50SP2ApplicationMessageReport(ApplicationMessageReport{msg}, sessionID)
	case "BZ":
		return router.OnFIX50SP2OrderMassActionReport(OrderMassActionReport{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP2OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP2UserNotification(UserNotification{msg}, sessionID)
	case "CC":
		return router.OnFIX50SP2StreamAssignmentRequest(StreamAssignmentRequest{msg}, sessionID)
	case "CD":
		return router.OnFIX50SP2StreamAssignmentReport(StreamAssignmentReport{msg}, sessionID)
	case "CE":
		return router.OnFIX50SP2StreamAssignmentReportACK(StreamAssignmentReportACK{msg}, sessionID)
	case "CF":
		return router.OnFIX50SP2PartyDetailsListRequest(PartyDetailsListRequest{msg}, sessionID)
	case "CG":
		return router.OnFIX50SP2PartyDetailsListReport(PartyDetailsListReport{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50SP2IOI(msg IOI, sessionID session.ID) reject.MessageReject
	OnFIX50SP2Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP2News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX50SP2Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MassQuote(msg MassQuote, sessionID session.ID) reject.MessageReject
	OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP2BidRequest(msg BidRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2BidResponse(msg BidResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionID session.ID) reject.MessageReject
	OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityList(msg SecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP2RFQRequest(msg RFQRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2Confirmation(msg Confirmation, sessionID session.ID) reject.MessageReject
	OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionID session.ID) reject.MessageReject
	OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2PositionReport(msg PositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AllocationReport(msg AllocationReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralReport(msg CollateralReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2UserRequest(msg UserRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2UserResponse(msg UserResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionID session.ID) reject.MessageReject
	OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2UserNotification(msg UserNotification, sessionID session.ID) reject.MessageReject
	OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionID session.ID) reject.MessageReject
	OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionID session.ID) reject.MessageReject
}
type FIX50SP2MessageCracker struct{}

func (c *FIX50SP2MessageCracker) OnFIX50SP2IOI(msg IOI, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuote(msg MassQuote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidRequest(msg BidRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidResponse(msg BidResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityList(msg SecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RFQRequest(msg RFQRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Confirmation(msg Confirmation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionReport(msg PositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReport(msg AllocationReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralReport(msg CollateralReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserRequest(msg UserRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserResponse(msg UserResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserNotification(msg UserNotification, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
