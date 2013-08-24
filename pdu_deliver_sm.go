package smpp34

import (
	"bytes"
)

var (
	reqDSMFields = []string{
		SERVICE_TYPE,
		SOURCE_ADDR_TON,
		SOURCE_ADDR_NPI,
		SOURCE_ADDR,
		DEST_ADDR_TON,
		DEST_ADDR_NPI,
		DESTINATION_ADDR,
		ESM_CLASS,
		PROTOCOL_ID,
		PRIORITY_FLAG,
		SCHEDULE_DELIVERY_TIME,
		VALIDITY_PERIOD,
		REGISTERED_DELIVERY,
		REPLACE_IF_PRESENT_FLAG,
		DATA_CODING,
		SM_DEFAULT_MSG_ID,
		SM_LENGTH,
		SHORT_MESSAGE,
	}
)

type DeliverSm struct {
	*Header
	mandatoryFields map[string]Field
	tlvFields       []*TLVField
}

func NewDeliverSm(hdr *Header, b []byte) (*DeliverSm, error) {
	r := bytes.NewBuffer(b)

	fields, tlvs, err := create_pdu_fields(reqDSMFields, r)

	if err != nil {
		return nil, err
	}

	d := &DeliverSm{hdr, fields, tlvs}

	return d, nil
}

func (d *DeliverSm) GetField(f string) Field {
	return d.mandatoryFields[f]
}

func (d *DeliverSm) Fields() map[string]Field {
	return d.mandatoryFields
}

func (d *DeliverSm) MandatoryFieldsList() []string {
	return reqDSMFields
}

func (d *DeliverSm) GetHeader() *Header {
	return d.Header
}

func (d *DeliverSm) TLVFields() []*TLVField {
	return d.tlvFields
}

func (d *DeliverSm) writeFields() []byte {
	b := []byte{}

	for _, i := range d.MandatoryFieldsList() {
		v := d.mandatoryFields[i].ByteArray()
		b = append(b, v...)
	}

	return b
}

func (d *DeliverSm) writeTLVFields() []byte {
	b := []byte{}

	for _, v := range d.tlvFields {
		b = append(b, v.Writer()...)
	}

	return b
}

func (d *DeliverSm) Writer() []byte {
	b := append(d.writeFields(), d.writeTLVFields()...)
	h := packUi32(uint32(len(b) + 16))
	h = append(h, packUi32(DELIVER_SM)...)
	h = append(h, packUi32(d.Header.Status)...)
	h = append(h, packUi32(d.Header.Sequence)...)

	return append(h, b...)
}
