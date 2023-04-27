package util

type RecordNotFound struct{}

func (m *RecordNotFound) Error() string {
	return "record_not_found"
}

type NotPublicIp struct{}

func (m *NotPublicIp) Error() string {
	return "Ip Address not public"
}
