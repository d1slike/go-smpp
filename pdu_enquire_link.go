package smpp34

var (
	reqELFields = []string{}
)

type EnquireLink struct {
	*Header
	mandatoryFields map[string]Field
	tlvFields       []*TLVField
}

func NewEnquireLink(hdr *Header) (*EnquireLink, error) {
	s := &EnquireLink{Header: hdr}

	return s, nil
}

func (s *EnquireLink) GetField(f string) Field {
	return nil
}

func (s *EnquireLink) Fields() map[string]Field {
	return s.mandatoryFields
}

func (s *EnquireLink) MandatoryFieldsList() []string {
	return reqELFields
}

func (s *EnquireLink) GetHeader() *Header {
	return s.Header
}

func (s *EnquireLink) TLVFields() []*TLVField {
	return s.tlvFields
}

func (s *EnquireLink) writeFields() []byte {
	return []byte{}
}

func (s *EnquireLink) Writer() []byte {
	b := s.writeFields()
	h := packUi32(uint32(len(b) + 16))
	h = append(h, packUi32(ENQUIRE_LINK)...)
	h = append(h, packUi32(s.Header.Status)...)
	h = append(h, packUi32(s.Header.Sequence)...)
	return append(h, b...)
}
