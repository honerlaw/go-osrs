package io

/**
This was ripped from rune-source (and in turn hyperion?) and converted to go
 */

const (
	RATIO = 0x9e3779b9
	SIZE_LOG = 8
	SIZE = 1 << SIZE_LOG
	MASK = (SIZE - 1) << 2
)

type Isaac struct {
	count   uint32
	results []uint32
	memory  []uint32
	a, b, c uint32
}

func NewIsaac(seeds []uint32) *Isaac {
	var isaac = &Isaac{
		count:   0,
		results: make([]uint32, SIZE),
		memory:  make([]uint32, SIZE),
	}

	for i, seed := range seeds {
		isaac.results[i] = seed
	}

	isaac.init(true)

	return isaac
}

func (i *Isaac) NextValue() uint32 {
	i.count -= 1
	if i.count == 0 {
		i.isaac()
		i.count = SIZE - 1
	}
	return i.results[i.count]
}

/**
 * Generates 256 results.
 */
func (is *Isaac) isaac() {
	var j = SIZE / 2
	var i, x, y uint32
	is.c += 1
	is.b += is.c;
	for i = 0; i < SIZE/2; {
		x = is.memory[i];
		is.a ^= is.a << 13;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a >> 6;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a << 2;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a >> 16;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
	}

	for j = 0; j < SIZE/2; {
		x = is.memory[i];
		is.a ^= is.a << 13;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a >> 6;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a << 2;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
		x = is.memory[i];
		is.a ^= is.a >> 16;
		is.a += is.memory[j];
		j++
		is.memory[i] = is.memory[(x&MASK)>>2] + is.a + is.b;
		y = is.memory[i]
		is.results[i] = is.memory[((y>>SIZE_LOG)&MASK)>>2] + x;
		is.b = is.results[i]
		i++
	}
}

/**
 * Initialises the ISAAC.
 *
 * @param flag
 *            Flag indicating if we should perform a second pass.
 */
func (is *Isaac) init(flag bool) {
	var i, a, b, c, d, e, f, g, h uint32 = 0, RATIO, RATIO, RATIO, RATIO, RATIO, RATIO, RATIO, RATIO
	for i = 0; i < 4; i++ {
		a ^= b << 11;
		d += a;
		b += c;
		b ^= c >> 2;
		e += b;
		c += d;
		c ^= d << 8;
		f += c;
		d += e;
		d ^= e >> 16;
		g += d;
		e += f;
		e ^= f << 10;
		h += e;
		f += g;
		f ^= g >> 4;
		a += f;
		g += h;
		g ^= h << 8;
		b += g;
		h += a;
		h ^= a >> 9;
		c += h;
		a += b;
	}
	for i = 0; i < SIZE; i += 8 {
		if flag {
			a += is.results[i];
			b += is.results[i+1];
			c += is.results[i+2];
			d += is.results[i+3];
			e += is.results[i+4];
			f += is.results[i+5];
			g += is.results[i+6];
			h += is.results[i+7];
		}
		a ^= b << 11;
		d += a;
		b += c;
		b ^= c >> 2;
		e += b;
		c += d;
		c ^= d << 8;
		f += c;
		d += e;
		d ^= e >> 16;
		g += d;
		e += f;
		e ^= f << 10;
		h += e;
		f += g;
		f ^= g >> 4;
		a += f;
		g += h;
		g ^= h << 8;
		b += g;
		h += a;
		h ^= a >> 9;
		c += h;
		a += b;
		is.memory[i] = a;
		is.memory[i+1] = b;
		is.memory[i+2] = c;
		is.memory[i+3] = d;
		is.memory[i+4] = e;
		is.memory[i+5] = f;
		is.memory[i+6] = g;
		is.memory[i+7] = h;
	}
	if flag {
		for i = 0; i < SIZE; i += 8 {
			a += is.memory[i];
			b += is.memory[i+1];
			c += is.memory[i+2];
			d += is.memory[i+3];
			e += is.memory[i+4];
			f += is.memory[i+5];
			g += is.memory[i+6];
			h += is.memory[i+7];
			a ^= b << 11;
			d += a;
			b += c;
			b ^= c >> 2;
			e += b;
			c += d;
			c ^= d << 8;
			f += c;
			d += e;
			d ^= e >> 16;
			g += d;
			e += f;
			e ^= f << 10;
			h += e;
			f += g;
			f ^= g >> 4;
			a += f;
			g += h;
			g ^= h << 8;
			b += g;
			h += a;
			h ^= a >> 9;
			c += h;
			a += b;
			is.memory[i] = a;
			is.memory[i+1] = b;
			is.memory[i+2] = c;
			is.memory[i+3] = d;
			is.memory[i+4] = e;
			is.memory[i+5] = f;
			is.memory[i+6] = g;
			is.memory[i+7] = h;
		}
	}
	is.isaac();
	is.count = SIZE;
}
