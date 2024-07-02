package main


const (
	// Header Constants
	QR_QUERY = 0
	QR_RESPONSE = 1
	OPCODE_QUERY  = 0  	// Standard query
	OPCODE_IQUERY = 1  	// Inverse query
	OPCODE_STATUS = 2  	// Server status request
	RD = 1              // Recusion Desired
	AA = 0              // Authoritative Answer
	TC = 0              // Trucation
	RA = 0              // Recursion Available
	Z  = 0              // Reserved for future use. Must be zero in all queries and responses.
	
	// Header response code types
	RCODE_NO_ERROR = 0
	RCODE_FORMAT_ERROR = 1
	RCODE_SERVER_FAILURE = 2
	RCODE_NAME_ERROR = 3
	RCODE_NOT_IMPLEMENTED = 4
	RCODE_REFUSED = 5
)

type DNSHeader struct {
	ID      uint16 
	Flags   uint16 
	QDCount uint16  	// The number of entries in the question section.
	ANCount uint16  	// The number of resource records in the answer section.
	NSCount uint16  	// The number of name server resource records in the authority section.
	ARCount uint16  	// The number of resource records in the additional section.
}

type DNSQuestion struct {
	QName  []byte 		// A domain name represented as a sequence of labels.
	QType  uint16 		// The type of the query
	QClass uint16 		// The class of the query.
}

type DNSResourceRecord struct {
	Name     []byte
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    []byte
}

/*     
		DNS Query Format

	+---------------------+
    |        Header       | Always present
    +---------------------+
    |       Question      | The question for the name server
    +---------------------+
    |        Answer       | RRs answering the question
    +---------------------+
    |      Authority      | RRs pointing toward an authority
    +---------------------+
    |      Additional     | RRs holding additional information
    +---------------------+

*/

type DNSMessage struct {
	Header      DNSHeader
	Question    DNSQuestion
	Answer     []DNSResourceRecord
	Authority  []DNSResourceRecord
	Additional []DNSResourceRecord
}



