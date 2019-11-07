package runepeeker

import (
	"bufio"
	"io"

	"github.com/romainmenke/css/tokenizer/streampreprocessor"
)

type RuneReader interface {
	ReadRune() (rune, int, error)
}

type Peeker struct {
	reader         RuneReader
	representation []rune

	peekBuffer []rune
	peekSizes  []int
}

func New(r *bufio.Reader) *Peeker {
	return &Peeker{
		reader: streampreprocessor.New(r),
	}
}

func (p *Peeker) Representation() []rune {
	return p.representation
}

func (p *Peeker) ResetRepresentation() {
	if len(p.representation) > 0 {
		p.representation = p.representation[:0]
	}
}

func (p *Peeker) buffRunes() {
	for len(p.peekBuffer) < 12 {
		r, size, _ := p.reader.ReadRune()
		if size == 0 {
			break
		}

		p.peekBuffer = append(p.peekBuffer, r)
		p.peekSizes = append(p.peekSizes, size)
	}
}

func (p *Peeker) readRune() (rune, int, error) {
	var (
		r    rune
		size int
		err  error
	)

	if len(p.peekBuffer) > 0 {
		r, p.peekBuffer = p.peekBuffer[0], p.peekBuffer[1:]
		size, p.peekSizes = p.peekSizes[0], p.peekSizes[1:]
	} else {
		r, size, err = p.reader.ReadRune()
		if err != nil {
			return r, size, err
		}
	}

	if size == 0 {
		return 0, 0, io.EOF
	}

	return r, size, err
}

func (p *Peeker) unreadRune(r rune, size int) {
	p.peekBuffer = append(p.peekBuffer, 0)
	copy(p.peekBuffer[1:], p.peekBuffer)
	p.peekBuffer[0] = r

	p.peekSizes = append(p.peekSizes, 0)
	copy(p.peekSizes[1:], p.peekSizes)
	p.peekSizes[0] = size
}

func (p *Peeker) ReadRune() (rune, int, error) {
	r, size, err := p.readRune()
	if size == 0 {
		return r, size, err
	}

	p.representation = append(p.representation, r)

	return r, size, err
}

func (p *Peeker) PeekRunes(n int) ([]rune, int, error) {
	var totalSize = 0

	toPeek := n - len(p.peekBuffer)
	for i := 0; i < toPeek; i++ {
		r, size, err := p.reader.ReadRune()
		if err != nil {
			return nil, 0, err
		}

		totalSize += size
		p.peekBuffer = append(p.peekBuffer, r)
		p.peekSizes = append(p.peekSizes, size)
	}

	out := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, p.peekBuffer[i])
	}

	return p.peekBuffer, totalSize, nil
}
