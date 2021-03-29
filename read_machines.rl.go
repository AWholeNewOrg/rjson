// Code generated by script/generate-ragel-file read_machines.rl. DO NOT EDIT.

package rjson

func readNull(data []byte) (int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)

	const readNull_start int = 1
	const readNull_first_final int = 6
	const readNull_error int = 0

	const readNull_en_main int = 1

	{
		cs = readNull_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 13:
			goto st2
		case 32:
			goto st2
		case 110:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st2
		}
		goto tr0
	tr0:
		return p, errNotNull
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 13:
			goto st2
		case 32:
			goto st2
		case 110:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st2
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 117 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 108 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 108 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5:
				return p, errNotNull
			}
		}

	_out:
		{
		}
	}

	return p, nil
}

func readBool(data []byte) (bool, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var val bool

	const readBool_start int = 1
	const readBool_first_final int = 10
	const readBool_error int = 0

	const readBool_en_main int = 1

	{
		cs = readBool_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 10:
			goto st_case_10
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 11:
			goto st_case_11
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 13:
			goto st2
		case 32:
			goto st2
		case 102:
			goto st3
		case 116:
			goto st7
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st2
		}
		goto tr0
	tr0:
		return false, p, errNotBool
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 13:
			goto st2
		case 32:
			goto st2
		case 102:
			goto st3
		case 116:
			goto st7
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st2
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 97 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 108 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 115 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 101 {
			goto tr7
		}
		goto tr0
	tr7:
		val = false
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 114 {
			goto st8
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 117 {
			goto st9
		}
		goto tr0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 101 {
			goto tr10
		}
		goto tr0
	tr10:
		val = true
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5, 6, 7, 8, 9:
				return false, p, errNotBool
			}
		}

	_out:
		{
		}
	}

	return val, p, nil
}

func readFloat64(data []byte) (float64, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var start int
	var hasDecimal bool
	var hasExp bool

	const readFloat64_start int = 1
	const readFloat64_first_final int = 6
	const readFloat64_error int = 0

	const readFloat64_en_main int = 1

	{
		cs = readFloat64_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 6:
			goto st_case_6
		case 3:
			goto st_case_3
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 45:
			goto tr1
		case 48:
			goto tr2
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto tr0
	tr0:
		return 0, p, errInvalidNumber
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr1:
		start = p
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 48 {
			goto st6
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto tr0
	tr2:
		start = p
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 46:
			goto st3
		case 69:
			goto st4
		case 101:
			goto st4
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto tr0
	tr6:
		hasDecimal = true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 69:
			goto st4
		case 101:
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
	tr12:
		hasDecimal = true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 69:
			goto st4
		case 101:
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 43:
			goto st5
		case 45:
			goto st5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto tr0
	tr8:
		hasExp = true
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr13
		}
		goto st0
	tr13:
		hasExp = true
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr13
		}
		goto st0
	tr3:
		start = p
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 46:
			goto st3
		case 69:
			goto st4
		case 101:
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 46:
			goto st3
		case 69:
			goto st4
		case 101:
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5:
				return 0, p, errInvalidNumber
			}
		}

	_out:
		{
		}
	}

	n, err := readFloat64Helper(hasDecimal, hasExp, data[start:p])
	return n, p, err
}

func readInt64(data []byte) (int64, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var start int
	var neg bool

	const readInt64_start int = 1
	const readInt64_first_final int = 6
	const readInt64_error int = 0

	const readInt64_en_main int = 1

	{
		cs = readInt64_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 6:
			goto st_case_6
		case 4:
			goto st_case_4
		case 7:
			goto st_case_7
		case 5:
			goto st_case_5
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 45:
			goto tr1
		case 48:
			goto tr2
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto tr0
	tr0:
		return 0, p, errInvalidInt
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr1:
		neg = true
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 48 {
			goto tr2
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto tr0
	tr2:
		start = p
		start = p
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		goto tr4
	tr4:
		p--
		{
			p++
			cs = 6
			goto _out
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	tr3:
		start = p
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto tr5
	tr5:
		p--
		{
			p++
			cs = 7
			goto _out
		}
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto tr5
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5:

				if p == 0 {
					return 0, p, errInvalidInt
				}
				p--
				{
					p++
					cs = 0
					goto _out
				}

				return 0, p, errInvalidInt
			}
		}

	_out:
		{
		}
	}

	n, err := readInt64Helper(neg, data[start:p])
	return n, p, err
}

func readInt32(data []byte) (int32, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var start int
	var neg bool

	const readInt32_start int = 1
	const readInt32_first_final int = 6
	const readInt32_error int = 0

	const readInt32_en_main int = 1

	{
		cs = readInt32_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 6:
			goto st_case_6
		case 4:
			goto st_case_4
		case 7:
			goto st_case_7
		case 5:
			goto st_case_5
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 45:
			goto tr1
		case 48:
			goto tr2
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto tr0
	tr0:
		return 0, p, errInvalidInt
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr1:
		neg = true
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 48 {
			goto tr2
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto tr0
	tr2:
		start = p
		start = p
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		goto tr4
	tr4:
		p--
		{
			p++
			cs = 6
			goto _out
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	tr3:
		start = p
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto tr5
	tr5:
		p--
		{
			p++
			cs = 7
			goto _out
		}
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto tr5
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5:

				if p == 0 {
					return 0, p, errInvalidInt
				}
				p--
				{
					p++
					cs = 0
					goto _out
				}

				return 0, p, errInvalidInt
			}
		}

	_out:
		{
		}
	}

	n, err := readInt32Helper(neg, data[start:p])
	return n, p, err
}

func readUint64(data []byte) (uint64, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var start int
	var err error

	const readUint64_start int = 1
	const readUint64_first_final int = 5
	const readUint64_error int = 0

	const readUint64_en_main int = 1

	{
		cs = readUint64_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 5:
			goto st_case_5
		case 3:
			goto st_case_3
		case 6:
			goto st_case_6
		case 4:
			goto st_case_4
		}
		goto st_out
	st_case_1:
		if data[p] == 48 {
			goto tr1
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto tr0
	tr0:
		err = errInvalidUInt
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr1:
		start = p
		start = p
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		goto tr3
	tr3:
		p--
		{
			p++
			cs = 5
			goto _out
		}
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st0
	tr2:
		start = p
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr4
	tr4:
		p--
		{
			p++
			cs = 6
			goto _out
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr4
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4:

				if p == 0 {
					return 0, p, errInvalidUInt
				}
				p--
				{
					p++
					cs = 0
					goto _out
				}

				err = errInvalidUInt
			}
		}

	_out:
		{
		}
	}

	if err != nil {
		return 0, p, err
	}
	n, err := readUint64Helper(data[start:p])
	return n, p, err
}

func readUint32(data []byte) (uint32, int, error) {
	cs, p := 0, 0
	pe := len(data)
	eof := len(data)
	var start int
	var err error

	const readUint32_start int = 1
	const readUint32_first_final int = 5
	const readUint32_error int = 0

	const readUint32_en_main int = 1

	{
		cs = readUint32_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 5:
			goto st_case_5
		case 3:
			goto st_case_3
		case 6:
			goto st_case_6
		case 4:
			goto st_case_4
		}
		goto st_out
	st_case_1:
		if data[p] == 48 {
			goto tr1
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto tr0
	tr0:
		err = errInvalidUInt
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr1:
		start = p
		start = p
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		goto tr3
	tr3:
		p--
		{
			p++
			cs = 5
			goto _out
		}
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st0
	tr2:
		start = p
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr4
	tr4:
		p--
		{
			p++
			cs = 6
			goto _out
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 46:
			goto tr0
		case 69:
			goto tr0
		case 101:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr4
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4:

				if p == 0 {
					return 0, p, errInvalidUInt
				}
				p--
				{
					p++
					cs = 0
					goto _out
				}

				err = errInvalidUInt
			}
		}

	_out:
		{
		}
	}

	if err != nil {
		return 0, p, err
	}
	n, err := readUint32Helper(data[start:p])
	return n, p, err
}
