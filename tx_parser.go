package main

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hyperledger/fabric-protos-go/peer"
)

type ParsedProcessedTransaction struct {
	// *peer.ProcessedTransaction
	ValidationCode      int32                      //func (*peer.ProcessedTransaction).GetValidationCode() int32
	TransactionEnvelope *ParsedTransactionEnvelope //func (*peer.ProcessedTransaction).GetTransactionEnvelope() *common.Envelope
}

func (dpt *ParsedProcessedTransaction) DecodeProcessedTransaction(data []byte) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	processedTransaction := &peer.ProcessedTransaction{}
	processedTransaction.XXX_Unmarshal(data)

	decodedTransactionEnvelope := &ParsedTransactionEnvelope{}
	err := decodedTransactionEnvelope.DecodeTransactionEnvelope(processedTransaction.GetTransactionEnvelope())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}

	dpt.TransactionEnvelope = decodedTransactionEnvelope
	dpt.ValidationCode = processedTransaction.GetValidationCode()

	logger.Printf("DecodedProcessedTransaction: %+v\n", dpt)

	return nil
}

type ParsedTransactionEnvelope struct {
	// *common.Envelope
	Payload   *ParsedPayload //func (*common.Envelope).GetPayload() *common.Payload
	Signature []byte         //func (*common.Envelope).GetSignature() []byte
}

func (dte *ParsedTransactionEnvelope) DecodeTransactionEnvelope(envelope *common.Envelope) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	envPayload := &common.Payload{}
	envPayload.XXX_Unmarshal(envelope.GetPayload())

	decodedPayload := &ParsedPayload{}
	err := decodedPayload.DecodePayload(envPayload)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dte.Payload = decodedPayload

	dte.Signature = envelope.GetSignature()

	logger.Printf("DecodedTransactionEnvelope: %+v\n", dte)

	return nil
}

type ParsedPayload struct {
	// *common.Payload
	Header *ParsedHeader //func (*common.Payload).GetHeader() *common.Header
	Data   *ParsedData   //func (*common.Payload).GetData() []byte
}

func (dp *ParsedPayload) DecodePayload(payload *common.Payload) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedHeader := &ParsedHeader{}
	err := decodedHeader.DecodeHeader(payload.GetHeader())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dp.Header = decodedHeader

	payloadData := &peer.Transaction{}
	payloadData.XXX_Unmarshal(payload.GetData())

	decodedData := &ParsedData{}
	err = decodedData.DecodeData(payloadData)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dp.Data = decodedData

	logger.Printf("DecodedPayload: %+v\n", dp)

	return nil
}

type ParsedHeader struct {
	// *common.Header
	ChannelHeader   *ParsedChannelHeader   //func (*common.Header).GetChannelHeader() []byte
	SignatureHeader *ParsedSignatureHeader //func (*common.Header).GetSignatureHeader() []byte
}

func (dh *ParsedHeader) DecodeHeader(header *common.Header) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	channelHeader := &common.ChannelHeader{}
	channelHeader.XXX_Unmarshal(header.GetChannelHeader())

	decodedChannelHeader := &ParsedChannelHeader{}
	err := decodedChannelHeader.DecodeChannelHeader(channelHeader)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dh.ChannelHeader = decodedChannelHeader

	signatureHeader := &common.SignatureHeader{}
	signatureHeader.XXX_Unmarshal(header.GetSignatureHeader())
	decodedSignatureHeader := &ParsedSignatureHeader{}
	err = decodedSignatureHeader.DecodeSignatureHeader(signatureHeader)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dh.SignatureHeader = decodedSignatureHeader

	logger.Printf("DecodedHeader: %+v\n", dh)

	return nil
}

type ParsedChannelHeader struct {
	// *common.ChannelHeader
	Type        int32                  //func (*common.ChannelHeader).GetType() int32
	Version     int32                  //func (*common.ChannelHeader).GetVersion() int32
	Timestamp   *timestamppb.Timestamp //func (*common.ChannelHeader).GetTimestamp() *timestamppb.Timestamp
	ChannelId   string                 //func (*common.ChannelHeader).GetChannelId() string
	TxId        string                 //func (*common.ChannelHeader).GetTxId() string
	Epoch       uint64                 //func (*common.ChannelHeader).GetEpoch() uint64
	Extension   []byte                 //func (*common.ChannelHeader).GetExtension() []byte
	TlsCertHash []byte                 //func (*common.ChannelHeader).GetTlsCertHash() []byte
}

func (dch *ParsedChannelHeader) DecodeChannelHeader(channelHeader *common.ChannelHeader) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dch.Type = channelHeader.GetType()
	dch.Version = channelHeader.GetVersion()
	dch.Timestamp = channelHeader.GetTimestamp()
	dch.ChannelId = channelHeader.GetChannelId()
	dch.TxId = channelHeader.GetTxId()
	dch.Epoch = channelHeader.GetEpoch()
	dch.Extension = channelHeader.GetExtension()
	dch.TlsCertHash = channelHeader.GetTlsCertHash()

	logger.Printf("DecodedChannelHeader: %+v\n", dch)

	return nil
}

type ParsedSignatureHeader struct {
	// *common.SignatureHeader
	Creator *ParsedSerializedIdentity //func (*common.SignatureHeader).GetCreator() []byte
	Nonce   []byte                    //func (*common.SignatureHeader).GetNonce() []byte
}

func (dsh *ParsedSignatureHeader) DecodeSignatureHeader(signatureHeader *common.SignatureHeader) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	serializedIdentity := &msp.SerializedIdentity{}
	serializedIdentity.XXX_Unmarshal(signatureHeader.GetCreator())

	decodedSerializedIdentity := &ParsedSerializedIdentity{}
	err := decodedSerializedIdentity.DecodeSerializedIdentity(serializedIdentity)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dsh.Creator = decodedSerializedIdentity

	dsh.Nonce = signatureHeader.GetNonce()

	logger.Printf("DecodedSignatureHeader: %+v\n", dsh)

	return nil
}

type ParsedSerializedIdentity struct {
	// *msp.SerializedIdentity
	Mspid   string         //func (*msp.SerializedIdentity).GetMspid() string
	IdBytes *ParsedIdBytes //func (*msp.SerializedIdentity).GetIdBytes() []byte
}

func (dsi *ParsedSerializedIdentity) DecodeSerializedIdentity(serializedIdentity *msp.SerializedIdentity) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dsi.Mspid = serializedIdentity.GetMspid()

	decodedIdBytes := &ParsedIdBytes{}
	err := decodedIdBytes.DecodeIdBytes(serializedIdentity.GetIdBytes())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dsi.IdBytes = decodedIdBytes

	logger.Printf("DecodedSerializedIdentity: %+v\n", dsi)

	return nil
}

type ParsedIdBytes struct {
	// basic info of *x509.Certificate
	Signature             []byte
	SignatureAlgorithm    string            // (*x509.Certificate).SignatureAlgorithm() x509.SignatureAlgorithm
	PublicKeyAlgorithm    string            // (*x509.Certificate).PublicKeyAlgorithm() x509.PublicKeyAlgorithm
	PublicKey             interface{}       // (*x509.Certificate).PublicKey() interface{}
	Version               int               // (*x509.Certificate).Version() int
	SerialNumber          *big.Int          // (*x509.Certificate).SerialNumber() *math.big.Int
	Issuer                string            // (*x509.Certificate).Issuer() x509.Certificate.pkix.Name
	Subject               string            // (*x509.Certificate).Subject() x509.Certificate.pkix.Name
	NotBefore             time.Time         // (*x509.Certificate).NotBefore() time.Time
	NotAfter              time.Time         // (*x509.Certificate).NotAfter() time.Time
	KeyUsage              x509.KeyUsage     // (*x509.Certificate).KeyUsage() x509.KeyUsage
	Extensions            []ParsedExtension // (*x509.Certificate).Extensions() []pkix.Extension
	BasicConstraintsValid bool              // (*x509.Certificate).BasicConstraintsValid() bool
	IsCA                  bool              // (*x509.Certificate).IsCA() bool
	MaxPathLen            int               // (*x509.Certificate).MaxPathLen() int
	MaxPathLenZero        bool              // (*x509.Certificate).MaxPathLenZero() bool
	SubjectKeyId          []byte            // (*x509.Certificate).SubjectKeyId() []byte
	AuthorityKeyId        []byte            // (*x509.Certificate).AuthorityKeyId() []byte
	//...
}

func (dib *ParsedIdBytes) DecodeIdBytes(idBytes []byte) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	bl, _ := pem.Decode(idBytes)

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}

	dib.Signature = cert.Signature
	dib.SignatureAlgorithm = cert.SignatureAlgorithm.String()
	dib.PublicKeyAlgorithm = cert.PublicKeyAlgorithm.String()
	dib.PublicKey = cert.PublicKey
	dib.Version = cert.Version
	dib.SerialNumber = cert.SerialNumber
	dib.Issuer = cert.Issuer.String()
	dib.Subject = cert.Subject.String()
	dib.NotBefore = cert.NotBefore
	dib.NotAfter = cert.NotAfter
	dib.KeyUsage = cert.KeyUsage

	decodedExtensions := []ParsedExtension{}
	for _, extension := range cert.Extensions {
		decodedExtension := ParsedExtension{}
		decodedExtension.Id = extension.Id.String()
		decodedExtension.Critical = extension.Critical
		decodedExtension.Value = extension.Value
		decodedExtensions = append(decodedExtensions, decodedExtension)
	}
	dib.Extensions = decodedExtensions
	dib.BasicConstraintsValid = cert.BasicConstraintsValid
	dib.IsCA = cert.IsCA
	dib.MaxPathLen = cert.MaxPathLen
	dib.MaxPathLenZero = cert.MaxPathLenZero
	dib.SubjectKeyId = cert.SubjectKeyId
	dib.AuthorityKeyId = cert.AuthorityKeyId

	logger.Printf("DecodedIdBytes: %+v\n", dib)

	return nil
}

type ParsedExtension struct {
	// pkix.Extension
	Id       string //func (pkix.Extension).Id() asn1.ObjectIdentifier
	Critical bool   //func (pkix.Extension).Critical() bool
	Value    []byte //func (pkix.Extension).Value() []byte
}

type ParsedData struct {
	// *peer.Transaction
	Actions []*ParsedTransactionAction //func (*peer.Transaction).GetActions() []*peer.TransactionAction
}

func (dd *ParsedData) DecodeData(data *peer.Transaction) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedTransactionActions := []*ParsedTransactionAction{}
	for _, action := range data.GetActions() {
		decodedTransactionAction := &ParsedTransactionAction{}
		decodedTransactionAction.DecodeTransactionAction(action)
		decodedTransactionActions = append(decodedTransactionActions, decodedTransactionAction)
	}
	dd.Actions = decodedTransactionActions

	logger.Printf("DecodedData: %+v\n", dd)

	return nil
}

type ParsedTransactionAction struct {
	// *peer.TransactionAction
	Header  *ParsedTransactionActionHeader //func (*peer.TransactionAction).GetHeader() []byte
	Payload *ParsedChaincodeActionPayload  //func (*peer.TransactionAction).GetPayload() []byte
}

func (dta *ParsedTransactionAction) DecodeTransactionAction(action *peer.TransactionAction) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	transactionActionHeader := &common.SignatureHeader{}
	transactionActionHeader.XXX_Unmarshal(action.GetHeader())

	decodedTransactionActionHeader := &ParsedTransactionActionHeader{}
	err := decodedTransactionActionHeader.DecodeTransactionActionHeader(transactionActionHeader)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dta.Header = decodedTransactionActionHeader

	chaincodeActionPayload := &peer.ChaincodeActionPayload{}
	chaincodeActionPayload.XXX_Unmarshal(action.GetPayload())

	decodedChaincodeActionPayload := &ParsedChaincodeActionPayload{}
	err = decodedChaincodeActionPayload.DecodeChaincodeActionPayload(chaincodeActionPayload)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dta.Payload = decodedChaincodeActionPayload

	logger.Printf("DecodedTransactionAction: %+v\n", dta)

	return nil
}

type ParsedTransactionActionHeader struct {
	// *common.SignatureHeader
	Creator *ParsedSerializedIdentity //func (*common.SignatureHeader).GetCreator() []byte
	Nonce   []byte                    //func (*common.SignatureHeader).GetNonce() []byte
}

func (dsh *ParsedTransactionActionHeader) DecodeTransactionActionHeader(transactionActionHeader *common.SignatureHeader) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	serializedIdentity := &msp.SerializedIdentity{}
	serializedIdentity.XXX_Unmarshal(transactionActionHeader.GetCreator())

	decodedSerializedIdentity := &ParsedSerializedIdentity{}
	err := decodedSerializedIdentity.DecodeSerializedIdentity(serializedIdentity)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dsh.Creator = decodedSerializedIdentity

	dsh.Nonce = transactionActionHeader.GetNonce()

	logger.Printf("DecodedTransactionActionHeader: %+v\n", dsh)

	return nil
}

type ParsedChaincodeActionPayload struct {
	// *peer.ChaincodeActionPayload
	ChaincodeProposalPayload *ParsedChaincodeProposalPayload //func (*peer.ChaincodeActionPayload).GetChaincodeProposalPayload() []byte
	Action                   *ParsedChaincodeEndorsedAction  //func (*peer.ChaincodeActionPayload).GetAction() *peer.ChaincodeEndorsedAction
}

func (dcap *ParsedChaincodeActionPayload) DecodeChaincodeActionPayload(chaincodeActionPayload *peer.ChaincodeActionPayload) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	chaincodeProposalPayload := &peer.ChaincodeProposalPayload{}
	chaincodeProposalPayload.XXX_Unmarshal(chaincodeActionPayload.GetChaincodeProposalPayload())

	decodedChaincodeProposalPayload := &ParsedChaincodeProposalPayload{}
	err := decodedChaincodeProposalPayload.DecodeChaincodeProposalPayload(chaincodeProposalPayload)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcap.ChaincodeProposalPayload = decodedChaincodeProposalPayload

	decodedChaincodeEndorsedAction := &ParsedChaincodeEndorsedAction{}
	err = decodedChaincodeEndorsedAction.DecodeChaincodeEndorsedAction(chaincodeActionPayload.GetAction())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcap.Action = decodedChaincodeEndorsedAction

	logger.Printf("DecodedChaincodeActionPayload: %+v\n", dcap)

	return nil
}

type ParsedChaincodeProposalPayload struct {
	// *peer.ChaincodeProposalPayload
	Input        *ParsedChaincodeInvocationSpec //func (*peer.ChaincodeProposalPayload).GetInput() []byte
	TransientMap map[string][]byte              //func (*peer.ChaincodeProposalPayload).GetTransientMap() map[string][]byte
}

func (dcpp *ParsedChaincodeProposalPayload) DecodeChaincodeProposalPayload(chaincodeProposalPayload *peer.ChaincodeProposalPayload) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	chaincodeInvocationSpec := &peer.ChaincodeInvocationSpec{}
	chaincodeInvocationSpec.XXX_Unmarshal(chaincodeProposalPayload.GetInput())

	decodedChaincodeInvocationSpec := &ParsedChaincodeInvocationSpec{}
	err := decodedChaincodeInvocationSpec.DecodeChaincodeInvocationSpec(chaincodeInvocationSpec)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcpp.Input = decodedChaincodeInvocationSpec

	dcpp.TransientMap = chaincodeProposalPayload.GetTransientMap()

	logger.Printf("DecodedChaincodeProposalPayload: %+v\n", dcpp)

	return nil
}

type ParsedChaincodeInvocationSpec struct {
	// *peer.ChaincodeInvocationSpec
	ChaincodeSpec *ParsedChaincodeSpec //func (*peer.ChaincodeInvocationSpec).GetChaincodeSpec() *peer.ChaincodeSpec
}

func (dcis *ParsedChaincodeInvocationSpec) DecodeChaincodeInvocationSpec(chaincodeInvocationSpec *peer.ChaincodeInvocationSpec) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedChaincodeSpec := &ParsedChaincodeSpec{}
	err := decodedChaincodeSpec.DecodeChaincodeSpec(chaincodeInvocationSpec.GetChaincodeSpec())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcis.ChaincodeSpec = decodedChaincodeSpec

	logger.Printf("DecodedChaincodeInvocationSpec: %+v\n", dcis)

	return nil
}

type ParsedChaincodeSpec struct {
	// *peer.ChaincodeSpec
	Type        string                //func (*peer.ChaincodeSpec).GetType() ChaincodeSpec_Type
	ChaincodeId *ParsedChaincodeId    //func (*peer.ChaincodeSpec).GetChaincodeId() *peer.ChaincodeID
	Input       *ParsedChaincodeInput //func (*peer.ChaincodeSpec).GetInput() *peer.ChaincodeInput
	Timeout     int32                 //func (*peer.ChaincodeSpec).GetTimeout() int32
}

func (dcs *ParsedChaincodeSpec) DecodeChaincodeSpec(chaincodeSpec *peer.ChaincodeSpec) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dcs.Type = chaincodeSpec.GetType().String()

	decodedChaincodeId := &ParsedChaincodeId{}
	err := decodedChaincodeId.DecodeChaincodeId(chaincodeSpec.GetChaincodeId())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcs.ChaincodeId = decodedChaincodeId

	decodedChaincodeInput := &ParsedChaincodeInput{}
	err = decodedChaincodeInput.DecodeChaincodeInput(chaincodeSpec.GetInput())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcs.Input = decodedChaincodeInput

	dcs.Timeout = chaincodeSpec.GetTimeout()

	logger.Printf("DecodedChaincodeSpec: %+v\n", dcs)

	return nil
}

type ParsedChaincodeId struct {
	// *peer.ChaincodeID
	Name    string //func (*peer.ChaincodeID).GetName() string
	Version string //func (*peer.ChaincodeID).GetVersion() string
	Path    string //func (*peer.ChaincodeID).GetPath() string
}

func (dci *ParsedChaincodeId) DecodeChaincodeId(chaincodeId *peer.ChaincodeID) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dci.Name = chaincodeId.GetName()
	dci.Version = chaincodeId.GetVersion()
	dci.Path = chaincodeId.GetPath()

	logger.Printf("DecodedChaincodeId: %+v\n", dci)

	return nil
}

type ParsedChaincodeInput struct {
	// *peer.ChaincodeInput
	Args        *ParsedArgs       //func (*peer.ChaincodeInput).GetArgs() [][]byte
	Decorations map[string][]byte //func (*peer.ChaincodeInput).GetDecorations() map[string][]byte
	IsInit      bool              //func (*peer.ChaincodeInput).GetIsInit() bool
}

func (dci *ParsedChaincodeInput) DecodeChaincodeInput(chaincodeInput *peer.ChaincodeInput) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedArgs := &ParsedArgs{}
	err := decodedArgs.DecodeArgs(chaincodeInput.GetArgs())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dci.Args = decodedArgs

	dci.Decorations = chaincodeInput.GetDecorations()
	dci.IsInit = chaincodeInput.GetIsInit()

	logger.Printf("DecodedChaincodeInput: %+v\n", dci)

	return nil
}

type ParsedArgs struct {
	Args []string
}

func (da *ParsedArgs) DecodeArgs(args [][]byte) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedArgs := []string{}
	for _, arg := range args {
		decodedArgs = append(decodedArgs, string(arg))
	}
	da.Args = decodedArgs

	logger.Printf("DecodedArgs: %+v\n", da)

	return nil
}

type ParsedChaincodeEndorsedAction struct {
	// *peer.ChaincodeEndorsedAction
	ProposalResponsePayload *ParsedProposalResponsePayload //func (*peer.ChaincodeEndorsedAction).GetProposalResponsePayload() []byte
	Endorsements            []*ParsedEndorsement           //func (*peer.ChaincodeEndorsedAction).GetEndorsements() []*peer.Endorsement
}

func (dcea *ParsedChaincodeEndorsedAction) DecodeChaincodeEndorsedAction(chaincodeEndorsedAction *peer.ChaincodeEndorsedAction) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	proposalResponsePayload := &peer.ProposalResponsePayload{}
	proposalResponsePayload.XXX_Unmarshal(chaincodeEndorsedAction.GetProposalResponsePayload())

	decodedProposalResponsePayload := &ParsedProposalResponsePayload{}
	err := decodedProposalResponsePayload.DecodeProposalResponsePayload(proposalResponsePayload)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dcea.ProposalResponsePayload = decodedProposalResponsePayload

	decodedEndorsements := []*ParsedEndorsement{}
	for _, endorsement := range chaincodeEndorsedAction.GetEndorsements() {
		decodedEndorsement := &ParsedEndorsement{}
		decodedEndorsement.DecodeEndorsement(endorsement)
		decodedEndorsements = append(decodedEndorsements, decodedEndorsement)
	}
	dcea.Endorsements = decodedEndorsements

	logger.Printf("DecodedChaincodeEndorsedAction: %+v\n", dcea)

	return nil
}

type ParsedProposalResponsePayload struct {
	// *peer.ProposalResponsePayload
	ProposalHash []byte                 //func (*peer.ProposalResponsePayload).GetProposalHash() []byte
	Extension    *ParsedChaincodeAction //func (*peer.ProposalResponsePayload).GetExtension() []byte
}

func (dprp *ParsedProposalResponsePayload) DecodeProposalResponsePayload(proposalResponsePayload *peer.ProposalResponsePayload) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dprp.ProposalHash = proposalResponsePayload.GetProposalHash()

	chaincodeAction := &peer.ChaincodeAction{}
	chaincodeAction.XXX_Unmarshal(proposalResponsePayload.GetExtension())

	decodedChaincodeAction := &ParsedChaincodeAction{}
	err := decodedChaincodeAction.DecodeChaincodeAction(chaincodeAction)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dprp.Extension = decodedChaincodeAction

	logger.Printf("DecodedProposalResponsePayload: %+v\n", dprp)

	return nil
}

type ParsedChaincodeAction struct {
	// *peer.ChaincodeAction
	Results     *ParsedReadWriteSet   //func (*peer.ChaincodeAction).GetResults() []byte
	Events      *ParsedChaincodeEvent //func (*peer.ChaincodeAction).GetEvents() []byte
	Response    *ParsedResponse       //func (*peer.ChaincodeAction).GetResponse() *peer.Response
	ChaincodeId *ParsedChaincodeId    //func (*peer.ChaincodeAction).GetChaincodeId() *peer.ChaincodeID
}

func (dca *ParsedChaincodeAction) DecodeChaincodeAction(chaincodeAction *peer.ChaincodeAction) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	txReadWriteSet := &rwset.TxReadWriteSet{}
	txReadWriteSet.XXX_Unmarshal(chaincodeAction.GetResults())
	decodedReadWriteSet := &ParsedReadWriteSet{}
	err := decodedReadWriteSet.DecodeReadWriteSet(txReadWriteSet)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dca.Results = decodedReadWriteSet

	chaincodeEvent := &peer.ChaincodeEvent{}
	chaincodeEvent.XXX_Unmarshal(chaincodeAction.GetEvents())
	decodedChaincodeEvent := &ParsedChaincodeEvent{}
	err = decodedChaincodeEvent.DecodeChaincodeEvent(chaincodeEvent)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dca.Events = decodedChaincodeEvent

	decodedResponse := &ParsedResponse{}
	err = decodedResponse.DecodeResponse(chaincodeAction.GetResponse())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dca.Response = decodedResponse

	decodedChaincodeId := &ParsedChaincodeId{}
	err = decodedChaincodeId.DecodeChaincodeId(chaincodeAction.GetChaincodeId())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dca.ChaincodeId = decodedChaincodeId

	logger.Printf("DecodedChaincodeAction: %+v\n", dca)

	return nil
}

type ParsedReadWriteSet struct {
	// *rwset.TxReadWriteSet
	DataModel string                  //func (*rwset.TxReadWriteSet).GetDataModel() rwset.TxReadWriteSet_DataModel
	NsRwset   []*ParsedNsReadWriteSet //func (*rwset.TxReadWriteSet).GetNsRwset() []*rwset.NsReadWriteSet
}

func (drws *ParsedReadWriteSet) DecodeReadWriteSet(txReadWriteSet *rwset.TxReadWriteSet) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	drws.DataModel = txReadWriteSet.GetDataModel().String()

	decodedNsReadWriteSets := []*ParsedNsReadWriteSet{}
	for _, nsReadWriteSet := range txReadWriteSet.GetNsRwset() {
		decodedNsReadWriteSet := &ParsedNsReadWriteSet{}
		decodedNsReadWriteSet.DecodeNsReadWriteSet(nsReadWriteSet)
		decodedNsReadWriteSets = append(decodedNsReadWriteSets, decodedNsReadWriteSet)
	}
	drws.NsRwset = decodedNsReadWriteSets

	logger.Printf("DecodedReadWriteSet: %+v\n", drws)

	return nil
}

type ParsedNsReadWriteSet struct {
	// *rwset.NsReadWriteSet
	Namespace             string                                //func (*rwset.NsReadWriteSet).GetNamespace() string
	Rwset                 *ParsedKVRWSet                        //func (*rwset.NsReadWriteSet).GetRwset() []byte
	CollectionHashedRwset []*ParsedCollectionHashedReadWriteSet //func (*rwset.NsReadWriteSet).GetCollectionHashedRwset() []*rwset.CollectionHashedReadWriteSet
}

func (dnrws *ParsedNsReadWriteSet) DecodeNsReadWriteSet(nsReadWriteSet *rwset.NsReadWriteSet) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dnrws.Namespace = nsReadWriteSet.GetNamespace()

	kvRwset := &kvrwset.KVRWSet{}
	kvRwset.XXX_Unmarshal(nsReadWriteSet.GetRwset())
	decodedKVRWSet := &ParsedKVRWSet{}
	err := decodedKVRWSet.DecodeKVRWSet(kvRwset)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dnrws.Rwset = decodedKVRWSet

	decodedCollectionHashedReadWriteSets := []*ParsedCollectionHashedReadWriteSet{}
	for _, collectionHashedReadWriteSet := range nsReadWriteSet.GetCollectionHashedRwset() {
		decodedCollectionHashedReadWriteSet := &ParsedCollectionHashedReadWriteSet{}
		decodedCollectionHashedReadWriteSet.DecodeCollectionHashedReadWriteSet(collectionHashedReadWriteSet)
		decodedCollectionHashedReadWriteSets = append(decodedCollectionHashedReadWriteSets, decodedCollectionHashedReadWriteSet)
	}
	dnrws.CollectionHashedRwset = decodedCollectionHashedReadWriteSets

	logger.Printf("DecodedNsReadWriteSet: %+v\n", dnrws)

	return nil
}

type ParsedKVRWSet struct {
	// *kvrwset.KVRWSet
	Reads            []*ParsedKVRead          //func (*kvrwset.KVRWSet).GetReads() []*kvrwset.KVRead
	RangeQueriesInfo []*ParsedRangeQueryInfo  //func (*kvrwset.KVRWSet).GetRangeQueriesInfo() []*kvrwset.RangeQueryInfo
	Writes           []*ParsedKVWrite         //func (*kvrwset.KVRWSet).GetWrites() []*kvrwset.KVWrite
	MetadataWrites   []*ParsedKVMetadataWrite //func (*kvrwset.KVRWSet).GetMetadataWrite() []*kvrwset.KVMetadataWrite
}

func (dkrws *ParsedKVRWSet) DecodeKVRWSet(kvRwset *kvrwset.KVRWSet) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedKVReads := []*ParsedKVRead{}
	for _, kvRead := range kvRwset.GetReads() {
		decodedKVRead := &ParsedKVRead{}
		decodedKVRead.DecodeKVRead(kvRead)
		decodedKVReads = append(decodedKVReads, decodedKVRead)
	}
	dkrws.Reads = decodedKVReads

	decodedRangeQueryInfos := []*ParsedRangeQueryInfo{}
	for _, rangeQueryInfo := range kvRwset.GetRangeQueriesInfo() {
		decodedRangeQueryInfo := &ParsedRangeQueryInfo{}
		decodedRangeQueryInfo.DecodeRangeQueryInfo(rangeQueryInfo)
		decodedRangeQueryInfos = append(decodedRangeQueryInfos, decodedRangeQueryInfo)
	}
	dkrws.RangeQueriesInfo = decodedRangeQueryInfos

	decodedKVWrites := []*ParsedKVWrite{}
	for _, kvWrite := range kvRwset.GetWrites() {
		decodedKVWrite := &ParsedKVWrite{}
		decodedKVWrite.DecodeKVWrite(kvWrite)
		decodedKVWrites = append(decodedKVWrites, decodedKVWrite)
	}
	dkrws.Writes = decodedKVWrites

	decodedKVMetadataWrites := []*ParsedKVMetadataWrite{}
	for _, kvMetadataWrite := range kvRwset.GetMetadataWrites() {
		decodedKVMetadataWrite := &ParsedKVMetadataWrite{}
		decodedKVMetadataWrite.DecodeKVMetadataWrite(kvMetadataWrite)
		decodedKVMetadataWrites = append(decodedKVMetadataWrites, decodedKVMetadataWrite)
	}
	dkrws.MetadataWrites = decodedKVMetadataWrites

	logger.Printf("DecodedKVRWSet: %+v\n", dkrws)

	return nil
}

type ParsedKVRead struct {
	// *kvrwset.KVRead
	Key     string         //func (*kvrwset.KVRead).GetKey() string
	Version *ParsedVersion //func (*kvrwset.KVRead).GetVersion() *kvrwset.Version
}

func (dkr *ParsedKVRead) DecodeKVRead(kvRead *kvrwset.KVRead) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkr.Key = kvRead.GetKey()

	decodedVersion := &ParsedVersion{}
	err := decodedVersion.DecodeVersion(kvRead.GetVersion())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dkr.Version = decodedVersion

	logger.Printf("DecodedKVRead: %+v\n", dkr)

	return nil
}

type ParsedVersion struct {
	// *kvrwset.Version
	BlockNum uint64 //func (*kvrwset.Version).GetBlockNum() uint64
	TxNum    uint64 //func (*kvrwset.Version).GetTxNum() uint64
}

func (dv *ParsedVersion) DecodeVersion(version *kvrwset.Version) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dv.BlockNum = version.GetBlockNum()
	dv.TxNum = version.GetTxNum()

	logger.Printf("DecodedVersion: %+v\n", dv)

	return nil
}

type ParsedRangeQueryInfo struct {
	// *kvrwset.RangeQueryInfo
	StartKey     string //func (*kvrwset.RangeQueryInfo).GetStartKey() string
	EndKey       string //func (*kvrwset.RangeQueryInfo).GetEndKey() string
	ItrExhausted bool   //func (*kvrwset.RangeQueryInfo).GetItrExhausted() bool
	// ReadsInfo *Decoded //func (*kvrwset.RangeQueryInfo).GetReadsInfo() kvrwset.isRangeQueryInfo_ReadsInfo
}

func (drqi *ParsedRangeQueryInfo) DecodeRangeQueryInfo(rangeQueryInfo *kvrwset.RangeQueryInfo) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	drqi.StartKey = rangeQueryInfo.GetStartKey()
	drqi.EndKey = rangeQueryInfo.GetEndKey()
	drqi.ItrExhausted = rangeQueryInfo.GetItrExhausted()

	logger.Printf("DecodedRangeQueryInfo: %+v\n", drqi)

	return nil
}

type ParsedKVWrite struct {
	// *kvrwset.KVWrite
	Key      string //func (*kvrwset.KVWrite).GetKey() string
	IsDelete bool   //func (*kvrwset.KVWrite).GetIsDelete() bool
	Value    []byte //func (*kvrwset.KVWrite).GetValue() []byte
}

func (dkw *ParsedKVWrite) DecodeKVWrite(kvWrite *kvrwset.KVWrite) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkw.Key = kvWrite.GetKey()
	dkw.IsDelete = kvWrite.GetIsDelete()
	dkw.Value = kvWrite.GetValue()

	logger.Printf("DecodedKVWrite: %+v\n", dkw)

	return nil
}

type ParsedKVMetadataWrite struct {
	// *kvrwset.KVMetadataWrite
	Key     string                   //func (*kvrwset.KVMetadataWrite).GetKey() string
	Entries []*ParsedKVMetadataEntry //func (*kvrwset.KVMetadataWrite).GetEntries() []*kvrwset.KVMetadataEntry
}

func (dkmw *ParsedKVMetadataWrite) DecodeKVMetadataWrite(kvMetadataWrite *kvrwset.KVMetadataWrite) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkmw.Key = kvMetadataWrite.GetKey()

	decodedKVMetadataEntries := []*ParsedKVMetadataEntry{}
	for _, kvMetadataEntry := range kvMetadataWrite.GetEntries() {
		decodedKVMetadataEntry := &ParsedKVMetadataEntry{}
		decodedKVMetadataEntry.DecodeKVMetadataEntry(kvMetadataEntry)
		decodedKVMetadataEntries = append(decodedKVMetadataEntries, decodedKVMetadataEntry)
	}
	dkmw.Entries = decodedKVMetadataEntries

	logger.Printf("DecodedKVMetadataWrite: %+v\n", dkmw)

	return nil
}

type ParsedKVMetadataEntry struct {
	// *kvrwset.KVMetadataEntry
	Name  string //func (*kvrwset.KVMetadataEntry).GetName() string
	Value []byte //func (*kvrwset.KVMetadataEntry).GetValue() []byte
}

func (dkme *ParsedKVMetadataEntry) DecodeKVMetadataEntry(kvMetadataEntry *kvrwset.KVMetadataEntry) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkme.Name = kvMetadataEntry.GetName()
	dkme.Value = kvMetadataEntry.GetValue()

	logger.Printf("DecodedKVMetadataEntry: %+v\n", dkme)

	return nil
}

type ParsedCollectionHashedReadWriteSet struct {
	// *rwset.CollectionHashedReadWriteSet
	CollectionName string             //func (*rwset.CollectionHashedReadWriteSet).GetCollectionName() string
	HashedRwset    *ParsedHashedRWSet //func (*rwset.CollectionHashedReadWriteSet).GetHashedRwset() []byte
	PvtRwsetHash   []byte             //func (*rwset.CollectionHashedReadWriteSet).GetPvtRwsetHash() []byte
}

func (dchrw *ParsedCollectionHashedReadWriteSet) DecodeCollectionHashedReadWriteSet(collectionHashedReadWriteSet *rwset.CollectionHashedReadWriteSet) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dchrw.CollectionName = collectionHashedReadWriteSet.GetCollectionName()

	hashedRwset := &kvrwset.HashedRWSet{}
	hashedRwset.XXX_Unmarshal(collectionHashedReadWriteSet.GetHashedRwset())
	decodedHashedRWSet := &ParsedHashedRWSet{}
	err := decodedHashedRWSet.DecodeHashedRWSet(hashedRwset)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dchrw.HashedRwset = decodedHashedRWSet

	dchrw.PvtRwsetHash = collectionHashedReadWriteSet.GetPvtRwsetHash()

	logger.Printf("DecodedCollectionHashedReadWriteSet: %+v\n", dchrw)

	return nil
}

type ParsedHashedRWSet struct {
	// *kvrwset.HashedRWSet
	HashedReads    []*ParsedKVReadHash          //func (*kvrwset.HashedRWSet).GetHashedReads() []*kvrwset.KVReadHash
	HashedWrites   []*ParsedKVWriteHash         //func (*kvrwset.HashedRWSet).GetHashedWrites() []*kvrwset.KVWriteHash
	MetadataWrites []*ParsedKVMetadataWriteHash //func (*kvrwset.HashedRWSet).GetMetadataWrites() []*kvrwset.KVMetadataWriteHash
}

func (dhrws *ParsedHashedRWSet) DecodeHashedRWSet(hashedRWSet *kvrwset.HashedRWSet) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	decodedKVReadHashes := []*ParsedKVReadHash{}
	for _, kvReadHash := range hashedRWSet.GetHashedReads() {
		decodedKVReadHash := &ParsedKVReadHash{}
		decodedKVReadHash.DecodeKVReadHash(kvReadHash)
		decodedKVReadHashes = append(decodedKVReadHashes, decodedKVReadHash)
	}
	dhrws.HashedReads = decodedKVReadHashes

	decodedKVWriteHashes := []*ParsedKVWriteHash{}
	for _, kvWriteHash := range hashedRWSet.GetHashedWrites() {
		decodedKVWriteHash := &ParsedKVWriteHash{}
		decodedKVWriteHash.DecodeKVWriteHash(kvWriteHash)
		decodedKVWriteHashes = append(decodedKVWriteHashes, decodedKVWriteHash)
	}
	dhrws.HashedWrites = decodedKVWriteHashes

	decodedKVMetadataWriteHashes := []*ParsedKVMetadataWriteHash{}
	for _, kvMetadataWriteHash := range hashedRWSet.GetMetadataWrites() {
		decodedKVMetadataWriteHash := &ParsedKVMetadataWriteHash{}
		decodedKVMetadataWriteHash.DecodeKVMetadataWriteHash(kvMetadataWriteHash)
		decodedKVMetadataWriteHashes = append(decodedKVMetadataWriteHashes, decodedKVMetadataWriteHash)
	}
	dhrws.MetadataWrites = decodedKVMetadataWriteHashes

	logger.Printf("DecodedHashedRWSet: %+v\n", dhrws)

	return nil
}

type ParsedKVReadHash struct {
	// *kvrwset.KVReadHash
	KeyHash []byte         //func (*kvrwset.KVReadHash).GetKeyHash() []byte
	Version *ParsedVersion //func (*kvrwset.KVReadHash).GetVersion() *kvrwset.Version
}

func (dkrh *ParsedKVReadHash) DecodeKVReadHash(kvReadHash *kvrwset.KVReadHash) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkrh.KeyHash = kvReadHash.GetKeyHash()

	decodedVersion := &ParsedVersion{}
	err := decodedVersion.DecodeVersion(kvReadHash.GetVersion())
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	dkrh.Version = decodedVersion

	logger.Printf("DecodedKVReadHash: %+v\n", dkrh)

	return nil
}

type ParsedKVWriteHash struct {
	// *kvrwset.KVWriteHash
	KeyHash   []byte //func (*kvrwset.KVWriteHash).GetKeyHash() []byte
	IsDelete  bool   //func (*kvrwset.KVWriteHash).GetIsDelete() bool
	ValueHash []byte //func (*kvrwset.KVWriteHash).GetValueHash() []byte
	IsPurge   bool   //func (*kvrwset.KVWriteHash).GetIsPurge() bool
}

func (dkwh *ParsedKVWriteHash) DecodeKVWriteHash(kvWriteHash *kvrwset.KVWriteHash) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkwh.KeyHash = kvWriteHash.GetKeyHash()
	dkwh.IsDelete = kvWriteHash.GetIsDelete()
	dkwh.ValueHash = kvWriteHash.GetValueHash()
	dkwh.IsPurge = kvWriteHash.GetIsPurge()

	logger.Printf("DecodedKVWriteHash: %+v\n", dkwh)

	return nil
}

type ParsedKVMetadataWriteHash struct {
	// *kvrwset.KVMetadataWriteHash
	KeyHash []byte                   //func (*kvrwset.KVMetadataWriteHash).GetKeyHash() []byte
	Entries []*ParsedKVMetadataEntry //func (*kvrwset.KVMetadataWriteHash).GetEntries() []*kvrwset.KVMetadataEntry
}

func (dkmwh *ParsedKVMetadataWriteHash) DecodeKVMetadataWriteHash(kvMetadataWriteHash *kvrwset.KVMetadataWriteHash) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dkmwh.KeyHash = kvMetadataWriteHash.GetKeyHash()

	decodedKVMetadataEntries := []*ParsedKVMetadataEntry{}
	for _, kvMetadataEntry := range kvMetadataWriteHash.GetEntries() {
		decodedKVMetadataEntry := &ParsedKVMetadataEntry{}
		decodedKVMetadataEntry.DecodeKVMetadataEntry(kvMetadataEntry)
		decodedKVMetadataEntries = append(decodedKVMetadataEntries, decodedKVMetadataEntry)
	}
	dkmwh.Entries = decodedKVMetadataEntries

	logger.Printf("DecodedKVMetadataWriteHash: %+v\n", dkmwh)

	return nil
}

type ParsedChaincodeEvent struct {
	// *peer.ChaincodeEvent
	ChaincodeId string //func (*peer.ChaincodeEvent).GetChaincodeId() string
	TxId        string //func (*peer.ChaincodeEvent).GetTxId() string
	EventName   string //func (*peer.ChaincodeEvent).GetEventName() string
	Payload     []byte //func (*peer.ChaincodeEvent).GetPayload() []byte
}

func (dce *ParsedChaincodeEvent) DecodeChaincodeEvent(chaincodeEvent *peer.ChaincodeEvent) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dce.ChaincodeId = chaincodeEvent.GetChaincodeId()
	dce.TxId = chaincodeEvent.GetTxId()
	dce.EventName = chaincodeEvent.GetEventName()
	dce.Payload = chaincodeEvent.GetPayload()

	logger.Printf("DecodedChaincodeEvent: %+v\n", dce)

	return nil
}

type ParsedResponse struct {
	// *peer.Response
	Status  int32  //func (*peer.Response).GetStatus() int32
	Message string //func (*peer.Response).GetMessage() string
	Payload []byte //func (*peer.Response).GetPayload() []byte
}

func (dr *ParsedResponse) DecodeResponse(response *peer.Response) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	dr.Status = response.GetStatus()
	dr.Message = response.GetMessage()
	dr.Payload = response.GetPayload()

	logger.Printf("DecodedResponse: %+v\n", dr)

	return nil
}

type ParsedEndorsement struct {
	// *peer.Endorsement
	Endorser  *ParsedSerializedIdentity //func (*peer.Endorsement).GetEndorser() []byte
	Signature []byte                    //func (*peer.Endorsement).GetSignature() []byte
}

func (de *ParsedEndorsement) DecodeEndorsement(endorsement *peer.Endorsement) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	serializedIdentity := &msp.SerializedIdentity{}
	serializedIdentity.XXX_Unmarshal(endorsement.GetEndorser())

	decodedSerializedIdentity := &ParsedSerializedIdentity{}
	err := decodedSerializedIdentity.DecodeSerializedIdentity(serializedIdentity)
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		return err
	}
	de.Endorser = decodedSerializedIdentity

	de.Signature = endorsement.GetSignature()

	logger.Printf("DecodedEndorsement: %+v\n", de)

	return nil
}
