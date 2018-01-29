package types

import (
	"encoding/xml"
	"time"

	dsigtypes "github.com/russellhaering/goxmldsig/types"
)

type EntityDescriptor struct {
	XMLName    xml.Name  `xml:"urn:oasis:names:tc:SAML:2.0:metadata EntityDescriptor"`
	ValidUntil time.Time `xml:"validUntil,attr"`
	// SAML 2.0 8.3.6 Entity Identifier could be used to represent issuer
	EntityID         string           `xml:"entityID,attr"`
	SPSSODescriptor  SPSSODescriptor  `xml:"SPSSODescriptor"`
	IDPSSODescriptor IDPSSODescriptor `xml:"IDPSSODescriptor,omitempty"`
	Organization     *Organization    `xml:"urn:oasis:names:tc:SAML:2.0:metadata Organization,omitempty"`
	ContactPerson    *ContactPerson   `xml:"urn:oasis:names:tc:SAML:2.0:metadata ContactPerson,omitempty"`
	DigestMethod     []DigestMethod   `xml:"urn:oasis:names:tc:SAML:metadata:algsupport DigestMethod"`
}

type Endpoint struct {
	Binding          string `xml:"Binding,attr"`
	Location         string `xml:"Location,attr"`
	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

type IndexedEndpoint struct {
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
	Index    int    `xml:"index,attr"`
}

type SPSSODescriptor struct {
	XMLName                    xml.Name          `xml:"urn:oasis:names:tc:SAML:2.0:metadata SPSSODescriptor"`
	AuthnRequestsSigned        bool              `xml:"AuthnRequestsSigned,attr"`
	WantAssertionsSigned       bool              `xml:"WantAssertionsSigned,attr"`
	ProtocolSupportEnumeration string            `xml:"protocolSupportEnumeration,attr"`
	KeyDescriptors             []KeyDescriptor   `xml:"KeyDescriptor"`
	SingleLogoutServices       []Endpoint        `xml:"SingleLogoutService"`
	NameIDFormat               string            `xml:"NameIDFormat,omitempty"`
	AssertionConsumerServices  []IndexedEndpoint `xml:"AssertionConsumerService"`
}

type IDPSSODescriptor struct {
	XMLName                 xml.Name              `xml:"urn:oasis:names:tc:SAML:2.0:metadata IDPSSODescriptor"`
	WantAuthnRequestsSigned bool                  `xml:"WantAuthnRequestsSigned,attr"`
	KeyDescriptors          []KeyDescriptor       `xml:"KeyDescriptor"`
	NameIDFormats           []NameIDFormat        `xml:"NameIDFormat"`
	SingleSignOnServices    []SingleSignOnService `xml:"SingleSignOnService"`
	Attributes              []Attribute           `xml:"Attribute"`
}

type KeyDescriptor struct {
	XMLName           xml.Name           `xml:"urn:oasis:names:tc:SAML:2.0:metadata KeyDescriptor"`
	Use               string             `xml:"use,attr"`
	KeyInfo           dsigtypes.KeyInfo  `xml:"KeyInfo"`
	EncryptionMethods []EncryptionMethod `xml:"EncryptionMethod"`
}

type NameIDFormat struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:metadata NameIDFormat"`
	Value   string   `xml:",chardata"`
}

type SingleSignOnService struct {
	XMLName  xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:metadata SingleSignOnService"`
	Binding  string   `xml:"Binding,attr"`
	Location string   `xml:"Location,attr"`
}

// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.3.2.1
type Organization struct {
	OrganizationNames        []LocalizedName `xml:"OrganizationName"`
	OrganizationDisplayNames []LocalizedName `xml:"OrganizationDisplayName"`
	OrganizationURLs         []LocalizedURI  `xml:"OrganizationURL"`
}

// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.2.4
type LocalizedName struct {
	Lang  string `xml:"xml lang,attr"`
	Value string `xml:",chardata"`
}

// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.2.5
type LocalizedURI struct {
	Lang  string `xml:"xml lang,attr"`
	Value string `xml:",chardata"`
}

// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.3.2.2
type ContactPerson struct {
	ContactType      string `xml:"contactType,attr"`
	Company          string
	GivenName        string
	SurName          string
	EmailAddresses   []string `xml:"EmailAddress"`
	TelephoneNumbers []string `xml:"TelephoneNumber"`
}
