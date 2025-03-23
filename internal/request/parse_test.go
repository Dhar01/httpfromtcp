package request

import "io"

type chunkReader struct {
	data            string // the complete data to be read
	numBytesPerRead int    // how many bytes to return per Read call
	pos             int    // current position in data
}

// Read reads up to len(p) or numBytesPerRead bytes from the string per call
// its useful for simulating reading a variable number of bytes per chunk from a network connection.
func (cr *chunkReader) Read(p []byte) (n int, err error) {
	// check if we've reached the end of the data
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}

	// Calculate where this read should end
	// want to read at most numBytesPerRead bytes
	endIndex := cr.pos + cr.numBytesPerRead

	// if endIndex would go beyond the end of data, adjust it
	if endIndex > len(cr.data) {
		endIndex = len(cr.data)
	}

	// copy the data from our string to the provided byte slice
	// returns the number of bytes actually copied
	n = copy(p, []byte(cr.data[cr.pos:endIndex]))

	// update our position for the next read
	cr.pos += n

	// ensure we don't read more than numBytesPerRead bytes
	// this is a safety check in case the provided buffer is larger
	// than what we want to read per chunk
	if n > cr.numBytesPerRead {
		n = cr.numBytesPerRead
		// adjust position back to account for the bytes we didn't actually want to read
		cr.pos -= cr.numBytesPerRead
	}

	return n, nil
}
