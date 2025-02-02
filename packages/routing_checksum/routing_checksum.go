package routing_checksum

import "main.go/packages/interfaces"

type RoutingChecksum struct{}

func (r *RoutingChecksum) IsValid(routingNumber interfaces.DigitSequence) bool {
	if !routingNumber.HasCorrectLength() {
		return false
	}
	sequence := routingNumber.GetSequence()
	sum1 := sequence[0] + sequence[3] + sequence[6]
	sum2 := sequence[1] + sequence[4] + sequence[7]
	sum3 := sequence[2] + sequence[5] + sequence[8]
	return ((3*sum1)+(7*sum2)+ sum3)%10 == 0
