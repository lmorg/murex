package jsonlines

import "bytes"

func noQuote(slice []byte) bool {
	return !bytes.HasPrefix(slice, []byte{'"'}) || !bytes.HasSuffix(slice, []byte{'"'})
}

func noSquare(slice []byte) bool {
	return !bytes.HasPrefix(slice, []byte{'['}) || !bytes.HasSuffix(slice, []byte{']'})
}

func noCurly(slice []byte) bool {
	return !bytes.HasPrefix(slice, []byte{'{'}) || !bytes.HasSuffix(slice, []byte{'}'})
}
