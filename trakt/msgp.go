package trakt

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"time"

	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z Airs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Day"
	o = append(o, 0x83, 0xa3, 0x44, 0x61, 0x79)
	o = msgp.AppendString(o, z.Day)
	// string "Time"
	o = append(o, 0xa4, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Time)
	// string "Timezone"
	o = append(o, 0xa8, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Timezone)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Airs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Day":
			z.Day, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Time":
			z.Time, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Timezone":
			z.Timezone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Airs) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Day) + 5 + msgp.StringPrefixSize + len(z.Time) + 9 + msgp.StringPrefixSize + len(z.Timezone)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CalendarMovie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Released"
	o = append(o, 0x82, 0xa8, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64)
	o = msgp.AppendString(o, z.Released)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CalendarMovie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Released":
			z.Released, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CalendarMovie) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Released) + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CalendarShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "FirstAired"
	o = append(o, 0x83, 0xaa, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendString(o, z.FirstAired)
	// string "Episode"
	o = append(o, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.Episode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Episode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CalendarShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FirstAired":
			z.FirstAired, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Episode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Episode = nil
			} else {
				if z.Episode == nil {
					z.Episode = new(Episode)
				}
				bts, err = z.Episode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CalendarShow) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.FirstAired) + 8
	if z.Episode == nil {
		s += msgp.NilSize
	} else {
		s += z.Episode.Msgsize()
	}
	s += 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Code) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "DeviceCode"
	o = append(o, 0x85, 0xaa, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.DeviceCode)
	// string "UserCode"
	o = append(o, 0xa8, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.UserCode)
	// string "VerificationURL"
	o = append(o, 0xaf, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.VerificationURL)
	// string "ExpiresIn"
	o = append(o, 0xa9, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e)
	o = msgp.AppendInt(o, z.ExpiresIn)
	// string "Interval"
	o = append(o, 0xa8, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c)
	o = msgp.AppendInt(o, z.Interval)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Code) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DeviceCode":
			z.DeviceCode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "UserCode":
			z.UserCode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "VerificationURL":
			z.VerificationURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ExpiresIn":
			z.ExpiresIn, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Interval":
			z.Interval, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Code) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.DeviceCode) + 9 + msgp.StringPrefixSize + len(z.UserCode) + 16 + msgp.StringPrefixSize + len(z.VerificationURL) + 10 + msgp.IntSize + 9 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CollectedEpisode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "CollectedAt"
	o = append(o, 0x82, 0xab, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.CollectedAt)
	// string "Number"
	o = append(o, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.Number)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CollectedEpisode) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "CollectedAt":
			z.CollectedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Number":
			z.Number, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z CollectedEpisode) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.CollectedAt) + 7 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CollectedSeason) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Number"
	o = append(o, 0x82, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.Number)
	// string "Episodes"
	o = append(o, 0xa8, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Episodes)))
	for za0001 := range z.Episodes {
		if z.Episodes[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "CollectedAt"
			o = append(o, 0x82, 0xab, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74)
			o = msgp.AppendString(o, z.Episodes[za0001].CollectedAt)
			// string "Number"
			o = append(o, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
			o = msgp.AppendInt(o, z.Episodes[za0001].Number)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CollectedSeason) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Number":
			z.Number, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Episodes":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Episodes) >= int(zb0002) {
				z.Episodes = (z.Episodes)[:zb0002]
			} else {
				z.Episodes = make([]*CollectedEpisode, zb0002)
			}
			for za0001 := range z.Episodes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Episodes[za0001] = nil
				} else {
					if z.Episodes[za0001] == nil {
						z.Episodes[za0001] = new(CollectedEpisode)
					}
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "CollectedAt":
							z.Episodes[za0001].CollectedAt, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Number":
							z.Episodes[za0001].Number, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CollectedSeason) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Episodes {
		if z.Episodes[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 12 + msgp.StringPrefixSize + len(z.Episodes[za0001].CollectedAt) + 7 + msgp.IntSize
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CollectionMovie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "CollectedAt"
	o = append(o, 0x82, 0xab, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.CollectedAt)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CollectionMovie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "CollectedAt":
			z.CollectedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CollectionMovie) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.CollectedAt) + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CollectionShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "CollectedAt"
	o = append(o, 0x83, 0xab, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.CollectedAt)
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Seasons"
	o = append(o, 0xa7, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons)))
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Seasons[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CollectionShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "CollectedAt":
			z.CollectedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Seasons":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seasons) >= int(zb0002) {
				z.Seasons = (z.Seasons)[:zb0002]
			} else {
				z.Seasons = make([]*CollectedSeason, zb0002)
			}
			for za0001 := range z.Seasons {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Seasons[za0001] = nil
				} else {
					if z.Seasons[za0001] == nil {
						z.Seasons[za0001] = new(CollectedSeason)
					}
					bts, err = z.Seasons[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CollectionShow) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.CollectedAt) + 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Seasons[za0001].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Episode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "Number"
	o = append(o, 0x8b, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.Number)
	// string "Season"
	o = append(o, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	o = msgp.AppendInt(o, z.Season)
	// string "Title"
	o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "Absolute"
	o = append(o, 0xa8, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65)
	o = msgp.AppendInt(o, z.Absolute)
	// string "FirstAired"
	o = append(o, 0xaa, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendString(o, z.FirstAired)
	// string "Translations"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Translations)))
	for za0001 := range z.Translations {
		o = msgp.AppendString(o, z.Translations[za0001])
	}
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat32(o, z.Rating)
	// string "Votes"
	o = append(o, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "IDs"
	o = append(o, 0xa3, 0x49, 0x44, 0x73)
	if z.IDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.IDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Episode) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Number":
			z.Number, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Season":
			z.Season, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Absolute":
			z.Absolute, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "FirstAired":
			z.FirstAired, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Translations":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Translations) >= int(zb0002) {
				z.Translations = (z.Translations)[:zb0002]
			} else {
				z.Translations = make([]string, zb0002)
			}
			for za0001 := range z.Translations {
				z.Translations[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "IDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.IDs = nil
			} else {
				if z.IDs == nil {
					z.IDs = new(IDs)
				}
				bts, err = z.IDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Episode) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 7 + msgp.IntSize + 6 + msgp.StringPrefixSize + len(z.Title) + 9 + msgp.StringPrefixSize + len(z.Overview) + 9 + msgp.IntSize + 11 + msgp.StringPrefixSize + len(z.FirstAired) + 13 + msgp.ArrayHeaderSize
	for za0001 := range z.Translations {
		s += msgp.StringPrefixSize + len(z.Translations[za0001])
	}
	s += 7 + msgp.Float32Size + 6 + msgp.IntSize + 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	s += 4
	if z.IDs == nil {
		s += msgp.NilSize
	} else {
		s += z.IDs.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EpisodeSearchResults) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 4
		// string "Type"
		o = append(o, 0x84, 0xa4, 0x54, 0x79, 0x70, 0x65)
		o = msgp.AppendString(o, z[za0001].Type)
		// string "Score"
		o = append(o, 0xa5, 0x53, 0x63, 0x6f, 0x72, 0x65)
		o, err = msgp.AppendIntf(o, z[za0001].Score)
		if err != nil {
			return
		}
		// string "Episode"
		o = append(o, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
		if z[za0001].Episode == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].Episode.MarshalMsg(o)
			if err != nil {
				return
			}
		}
		// string "Show"
		o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
		if z[za0001].Show == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].Show.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EpisodeSearchResults) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(EpisodeSearchResults, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch msgp.UnsafeString(field) {
			case "Type":
				(*z)[zb0001].Type, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "Score":
				(*z)[zb0001].Score, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			case "Episode":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].Episode = nil
				} else {
					if (*z)[zb0001].Episode == nil {
						(*z)[zb0001].Episode = new(Episode)
					}
					bts, err = (*z)[zb0001].Episode.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			case "Show":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].Show = nil
				} else {
					if (*z)[zb0001].Show == nil {
						(*z)[zb0001].Show = new(Show)
					}
					bts, err = (*z)[zb0001].Show.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EpisodeSearchResults) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + 5 + msgp.StringPrefixSize + len(z[zb0004].Type) + 6 + msgp.GuessSize(z[zb0004].Score) + 8
		if z[zb0004].Episode == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].Episode.Msgsize()
		}
		s += 5
		if z[zb0004].Show == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].Show.Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IDs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Trakt"
	o = append(o, 0x86, 0xa5, 0x54, 0x72, 0x61, 0x6b, 0x74)
	o = msgp.AppendInt(o, z.Trakt)
	// string "IMDB"
	o = append(o, 0xa4, 0x49, 0x4d, 0x44, 0x42)
	o = msgp.AppendString(o, z.IMDB)
	// string "TMDB"
	o = append(o, 0xa4, 0x54, 0x4d, 0x44, 0x42)
	o = msgp.AppendInt(o, z.TMDB)
	// string "TVDB"
	o = append(o, 0xa4, 0x54, 0x56, 0x44, 0x42)
	o = msgp.AppendInt(o, z.TVDB)
	// string "TVRage"
	o = append(o, 0xa6, 0x54, 0x56, 0x52, 0x61, 0x67, 0x65)
	o = msgp.AppendInt(o, z.TVRage)
	// string "Slug"
	o = append(o, 0xa4, 0x53, 0x6c, 0x75, 0x67)
	o = msgp.AppendString(o, z.Slug)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IDs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Trakt":
			z.Trakt, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "IMDB":
			z.IMDB, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TMDB":
			z.TMDB, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TVDB":
			z.TVDB, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TVRage":
			z.TVRage, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Slug":
			z.Slug, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *IDs) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.IMDB) + 5 + msgp.IntSize + 5 + msgp.IntSize + 7 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Slug)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Images) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "Poster"
	o = append(o, 0x89, 0xa6, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72)
	if z.Poster == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.Poster.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.Poster.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.Poster.Thumbnail)
	}
	// string "FanArt"
	o = append(o, 0xa6, 0x46, 0x61, 0x6e, 0x41, 0x72, 0x74)
	if z.FanArt == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.FanArt.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.FanArt.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.FanArt.Thumbnail)
	}
	// string "ScreenShot"
	o = append(o, 0xaa, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x68, 0x6f, 0x74)
	if z.ScreenShot == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.ScreenShot.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.ScreenShot.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.ScreenShot.Thumbnail)
	}
	// string "HeadShot"
	o = append(o, 0xa8, 0x48, 0x65, 0x61, 0x64, 0x53, 0x68, 0x6f, 0x74)
	if z.HeadShot == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.HeadShot.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.HeadShot.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.HeadShot.Thumbnail)
	}
	// string "Logo"
	o = append(o, 0xa4, 0x4c, 0x6f, 0x67, 0x6f)
	if z.Logo == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.Logo.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.Logo.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.Logo.Thumbnail)
	}
	// string "ClearArt"
	o = append(o, 0xa8, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x72, 0x74)
	if z.ClearArt == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.ClearArt.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.ClearArt.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.ClearArt.Thumbnail)
	}
	// string "Banner"
	o = append(o, 0xa6, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72)
	if z.Banner == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.Banner.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.Banner.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.Banner.Thumbnail)
	}
	// string "Thumbnail"
	o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
	if z.Thumbnail == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.Thumbnail.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.Thumbnail.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.Thumbnail.Thumbnail)
	}
	// string "Avatar"
	o = append(o, 0xa6, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72)
	if z.Avatar == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Full"
		o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
		o = msgp.AppendString(o, z.Avatar.Full)
		// string "Medium"
		o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
		o = msgp.AppendString(o, z.Avatar.Medium)
		// string "Thumbnail"
		o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
		o = msgp.AppendString(o, z.Avatar.Thumbnail)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Images) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Poster":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Poster = nil
			} else {
				if z.Poster == nil {
					z.Poster = new(Sizes)
				}
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.Poster.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.Poster.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.Poster.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "FanArt":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.FanArt = nil
			} else {
				if z.FanArt == nil {
					z.FanArt = new(Sizes)
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.FanArt.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.FanArt.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.FanArt.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "ScreenShot":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ScreenShot = nil
			} else {
				if z.ScreenShot == nil {
					z.ScreenShot = new(Sizes)
				}
				var zb0004 uint32
				zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0004 > 0 {
					zb0004--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.ScreenShot.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.ScreenShot.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.ScreenShot.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "HeadShot":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.HeadShot = nil
			} else {
				if z.HeadShot == nil {
					z.HeadShot = new(Sizes)
				}
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.HeadShot.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.HeadShot.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.HeadShot.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Logo":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Logo = nil
			} else {
				if z.Logo == nil {
					z.Logo = new(Sizes)
				}
				var zb0006 uint32
				zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0006 > 0 {
					zb0006--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.Logo.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.Logo.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.Logo.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "ClearArt":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ClearArt = nil
			} else {
				if z.ClearArt == nil {
					z.ClearArt = new(Sizes)
				}
				var zb0007 uint32
				zb0007, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0007 > 0 {
					zb0007--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.ClearArt.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.ClearArt.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.ClearArt.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Banner":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Banner = nil
			} else {
				if z.Banner == nil {
					z.Banner = new(Sizes)
				}
				var zb0008 uint32
				zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0008 > 0 {
					zb0008--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.Banner.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.Banner.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.Banner.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Thumbnail":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Thumbnail = nil
			} else {
				if z.Thumbnail == nil {
					z.Thumbnail = new(Sizes)
				}
				var zb0009 uint32
				zb0009, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0009 > 0 {
					zb0009--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.Thumbnail.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.Thumbnail.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.Thumbnail.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Avatar":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Avatar = nil
			} else {
				if z.Avatar == nil {
					z.Avatar = new(Sizes)
				}
				var zb0010 uint32
				zb0010, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0010 > 0 {
					zb0010--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Full":
						z.Avatar.Full, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Medium":
						z.Avatar.Medium, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Thumbnail":
						z.Avatar.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Images) Msgsize() (s int) {
	s = 1 + 7
	if z.Poster == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Poster.Full) + 7 + msgp.StringPrefixSize + len(z.Poster.Medium) + 10 + msgp.StringPrefixSize + len(z.Poster.Thumbnail)
	}
	s += 7
	if z.FanArt == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.FanArt.Full) + 7 + msgp.StringPrefixSize + len(z.FanArt.Medium) + 10 + msgp.StringPrefixSize + len(z.FanArt.Thumbnail)
	}
	s += 11
	if z.ScreenShot == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.ScreenShot.Full) + 7 + msgp.StringPrefixSize + len(z.ScreenShot.Medium) + 10 + msgp.StringPrefixSize + len(z.ScreenShot.Thumbnail)
	}
	s += 9
	if z.HeadShot == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.HeadShot.Full) + 7 + msgp.StringPrefixSize + len(z.HeadShot.Medium) + 10 + msgp.StringPrefixSize + len(z.HeadShot.Thumbnail)
	}
	s += 5
	if z.Logo == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Logo.Full) + 7 + msgp.StringPrefixSize + len(z.Logo.Medium) + 10 + msgp.StringPrefixSize + len(z.Logo.Thumbnail)
	}
	s += 9
	if z.ClearArt == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.ClearArt.Full) + 7 + msgp.StringPrefixSize + len(z.ClearArt.Medium) + 10 + msgp.StringPrefixSize + len(z.ClearArt.Thumbnail)
	}
	s += 7
	if z.Banner == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Banner.Full) + 7 + msgp.StringPrefixSize + len(z.Banner.Medium) + 10 + msgp.StringPrefixSize + len(z.Banner.Thumbnail)
	}
	s += 10
	if z.Thumbnail == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Thumbnail.Full) + 7 + msgp.StringPrefixSize + len(z.Thumbnail.Medium) + 10 + msgp.StringPrefixSize + len(z.Thumbnail.Thumbnail)
	}
	s += 7
	if z.Avatar == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Avatar.Full) + 7 + msgp.StringPrefixSize + len(z.Avatar.Medium) + 10 + msgp.StringPrefixSize + len(z.Avatar.Thumbnail)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *List) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "Name"
	o = append(o, 0x8d, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Description"
	o = append(o, 0xab, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Description)
	// string "Privacy"
	o = append(o, 0xa7, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79)
	o = msgp.AppendString(o, z.Privacy)
	// string "DisplayNumbers"
	o = append(o, 0xae, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73)
	o = msgp.AppendBool(o, z.DisplayNumbers)
	// string "AllowComments"
	o = append(o, 0xad, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendBool(o, z.AllowComments)
	// string "SortBy"
	o = append(o, 0xa6, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79)
	o = msgp.AppendString(o, z.SortBy)
	// string "SortHow"
	o = append(o, 0xa7, 0x53, 0x6f, 0x72, 0x74, 0x48, 0x6f, 0x77)
	o = msgp.AppendString(o, z.SortHow)
	// string "CreatedAt"
	o = append(o, 0xa9, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.CreatedAt)
	// string "UpdatedAt"
	o = append(o, 0xa9, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.UpdatedAt)
	// string "ItemCount"
	o = append(o, 0xa9, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.ItemCount)
	// string "CommentCount"
	o = append(o, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.CommentCount)
	// string "Likes"
	o = append(o, 0xa5, 0x4c, 0x69, 0x6b, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Likes)
	// string "IDs"
	o = append(o, 0xa3, 0x49, 0x44, 0x73)
	if z.IDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.IDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *List) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Description":
			z.Description, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Privacy":
			z.Privacy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "DisplayNumbers":
			z.DisplayNumbers, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "AllowComments":
			z.AllowComments, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "SortBy":
			z.SortBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SortHow":
			z.SortHow, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CreatedAt":
			z.CreatedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "UpdatedAt":
			z.UpdatedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ItemCount":
			z.ItemCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "CommentCount":
			z.CommentCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Likes":
			z.Likes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "IDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.IDs = nil
			} else {
				if z.IDs == nil {
					z.IDs = new(IDs)
				}
				bts, err = z.IDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *List) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 12 + msgp.StringPrefixSize + len(z.Description) + 8 + msgp.StringPrefixSize + len(z.Privacy) + 15 + msgp.BoolSize + 14 + msgp.BoolSize + 7 + msgp.StringPrefixSize + len(z.SortBy) + 8 + msgp.StringPrefixSize + len(z.SortHow) + 10 + msgp.StringPrefixSize + len(z.CreatedAt) + 10 + msgp.StringPrefixSize + len(z.UpdatedAt) + 10 + msgp.IntSize + 13 + msgp.IntSize + 6 + msgp.IntSize + 4
	if z.IDs == nil {
		s += msgp.NilSize
	} else {
		s += z.IDs.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ListItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Rank"
	o = append(o, 0x85, 0xa4, 0x52, 0x61, 0x6e, 0x6b)
	o = msgp.AppendInt(o, z.Rank)
	// string "ListedAt"
	o = append(o, 0xa8, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.ListedAt)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ListItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Rank":
			z.Rank, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ListedAt":
			z.ListedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ListItem) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 9 + msgp.StringPrefixSize + len(z.ListedAt) + 5 + msgp.StringPrefixSize + len(z.Type) + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	s += 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Movie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 14
	// string "Object"
	o = append(o, 0x8e, 0xa6, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74)
	o, err = z.Object.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Released"
	o = append(o, 0xa8, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64)
	o = msgp.AppendString(o, z.Released)
	// string "URL"
	o = append(o, 0xa3, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.URL)
	// string "Trailer"
	o = append(o, 0xa7, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72)
	o = msgp.AppendString(o, z.Trailer)
	// string "Runtime"
	o = append(o, 0xa7, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.Runtime)
	// string "TagLine"
	o = append(o, 0xa7, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65)
	o = msgp.AppendString(o, z.TagLine)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "Certification"
	o = append(o, 0xad, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Certification)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat32(o, z.Rating)
	// string "Votes"
	o = append(o, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Genres"
	o = append(o, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0001 := range z.Genres {
		o = msgp.AppendString(o, z.Genres[za0001])
	}
	// string "Language"
	o = append(o, 0xa8, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Language)
	// string "Translations"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Translations)))
	for za0002 := range z.Translations {
		o = msgp.AppendString(o, z.Translations[za0002])
	}
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Movie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Object":
			bts, err = z.Object.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Released":
			z.Released, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "URL":
			z.URL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Trailer":
			z.Trailer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Runtime":
			z.Runtime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TagLine":
			z.TagLine, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Certification":
			z.Certification, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Genres":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0002) {
				z.Genres = (z.Genres)[:zb0002]
			} else {
				z.Genres = make([]string, zb0002)
			}
			for za0001 := range z.Genres {
				z.Genres[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Language":
			z.Language, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Translations":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Translations) >= int(zb0003) {
				z.Translations = (z.Translations)[:zb0003]
			} else {
				z.Translations = make([]string, zb0003)
			}
			for za0002 := range z.Translations {
				z.Translations[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Movie) Msgsize() (s int) {
	s = 1 + 7 + z.Object.Msgsize() + 9 + msgp.StringPrefixSize + len(z.Released) + 4 + msgp.StringPrefixSize + len(z.URL) + 8 + msgp.StringPrefixSize + len(z.Trailer) + 8 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.TagLine) + 9 + msgp.StringPrefixSize + len(z.Overview) + 14 + msgp.StringPrefixSize + len(z.Certification) + 7 + msgp.Float32Size + 6 + msgp.IntSize + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Genres {
		s += msgp.StringPrefixSize + len(z.Genres[za0001])
	}
	s += 9 + msgp.StringPrefixSize + len(z.Language) + 13 + msgp.ArrayHeaderSize
	for za0002 := range z.Translations {
		s += msgp.StringPrefixSize + len(z.Translations[za0002])
	}
	s += 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MovieSearchResults) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 3
		// string "Type"
		o = append(o, 0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
		o = msgp.AppendString(o, z[za0001].Type)
		// string "Score"
		o = append(o, 0xa5, 0x53, 0x63, 0x6f, 0x72, 0x65)
		o, err = msgp.AppendIntf(o, z[za0001].Score)
		if err != nil {
			return
		}
		// string "Movie"
		o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
		if z[za0001].Movie == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].Movie.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MovieSearchResults) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(MovieSearchResults, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch msgp.UnsafeString(field) {
			case "Type":
				(*z)[zb0001].Type, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "Score":
				(*z)[zb0001].Score, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			case "Movie":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].Movie = nil
				} else {
					if (*z)[zb0001].Movie == nil {
						(*z)[zb0001].Movie = new(Movie)
					}
					bts, err = (*z)[zb0001].Movie.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MovieSearchResults) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + 5 + msgp.StringPrefixSize + len(z[zb0004].Type) + 6 + msgp.GuessSize(z[zb0004].Score) + 6
		if z[zb0004].Movie == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].Movie.Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Movies) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Watchers"
	o = append(o, 0x82, 0xa8, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	o = msgp.AppendInt(o, z.Watchers)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Movies) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Watchers":
			z.Watchers, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Movies) Msgsize() (s int) {
	s = 1 + 9 + msgp.IntSize + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Title"
	o = append(o, 0x83, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "Year"
	o = append(o, 0xa4, 0x59, 0x65, 0x61, 0x72)
	o = msgp.AppendInt(o, z.Year)
	// string "IDs"
	o = append(o, 0xa3, 0x49, 0x44, 0x73)
	if z.IDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.IDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Year":
			z.Year, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "IDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.IDs = nil
			} else {
				if z.IDs == nil {
					z.IDs = new(IDs)
				}
				bts, err = z.IDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Title) + 5 + msgp.IntSize + 4
	if z.IDs == nil {
		s += msgp.NilSize
	} else {
		s += z.IDs.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ProgressShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Episode"
	o = append(o, 0x82, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.Episode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Episode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ProgressShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Episode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Episode = nil
			} else {
				if z.Episode == nil {
					z.Episode = new(Episode)
				}
				bts, err = z.Episode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProgressShow) Msgsize() (s int) {
	s = 1 + 8
	if z.Episode == nil {
		s += msgp.NilSize
	} else {
		s += z.Episode.Msgsize()
	}
	s += 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Season) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "Number"
	o = append(o, 0x88, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.Number)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "EpisodeCount"
	o = append(o, 0xac, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.EpisodeCount)
	// string "AiredEpisodes"
	o = append(o, 0xad, 0x41, 0x69, 0x72, 0x65, 0x64, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendInt(o, z.AiredEpisodes)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat32(o, z.Rating)
	// string "Votes"
	o = append(o, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "IDs"
	o = append(o, 0xa3, 0x49, 0x44, 0x73)
	if z.IDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.IDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Season) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Number":
			z.Number, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeCount":
			z.EpisodeCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "AiredEpisodes":
			z.AiredEpisodes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "IDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.IDs = nil
			} else {
				if z.IDs == nil {
					z.IDs = new(IDs)
				}
				bts, err = z.IDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Season) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 9 + msgp.StringPrefixSize + len(z.Overview) + 13 + msgp.IntSize + 14 + msgp.IntSize + 7 + msgp.Float32Size + 6 + msgp.IntSize + 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	s += 4
	if z.IDs == nil {
		s += msgp.NilSize
	} else {
		s += z.IDs.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Show) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 18
	// string "Object"
	o = append(o, 0xde, 0x0, 0x12, 0xa6, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74)
	o, err = z.Object.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "FirstAired"
	o = append(o, 0xaa, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendString(o, z.FirstAired)
	// string "URL"
	o = append(o, 0xa3, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.URL)
	// string "Trailer"
	o = append(o, 0xa7, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72)
	o = msgp.AppendString(o, z.Trailer)
	// string "Runtime"
	o = append(o, 0xa7, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.Runtime)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "Certification"
	o = append(o, 0xad, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Certification)
	// string "Status"
	o = append(o, 0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "Network"
	o = append(o, 0xa7, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	o = msgp.AppendString(o, z.Network)
	// string "AiredEpisodes"
	o = append(o, 0xad, 0x41, 0x69, 0x72, 0x65, 0x64, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendInt(o, z.AiredEpisodes)
	// string "Airs"
	o = append(o, 0xa4, 0x41, 0x69, 0x72, 0x73)
	if z.Airs == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "Day"
		o = append(o, 0x83, 0xa3, 0x44, 0x61, 0x79)
		o = msgp.AppendString(o, z.Airs.Day)
		// string "Time"
		o = append(o, 0xa4, 0x54, 0x69, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Airs.Time)
		// string "Timezone"
		o = append(o, 0xa8, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65)
		o = msgp.AppendString(o, z.Airs.Timezone)
	}
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat32(o, z.Rating)
	// string "Votes"
	o = append(o, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Genres"
	o = append(o, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0001 := range z.Genres {
		o = msgp.AppendString(o, z.Genres[za0001])
	}
	// string "Country"
	o = append(o, 0xa7, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79)
	o = msgp.AppendString(o, z.Country)
	// string "Language"
	o = append(o, 0xa8, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Language)
	// string "Translations"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Translations)))
	for za0002 := range z.Translations {
		o = msgp.AppendString(o, z.Translations[za0002])
	}
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Show) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Object":
			bts, err = z.Object.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "FirstAired":
			z.FirstAired, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "URL":
			z.URL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Trailer":
			z.Trailer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Runtime":
			z.Runtime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Certification":
			z.Certification, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Network":
			z.Network, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AiredEpisodes":
			z.AiredEpisodes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Airs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Airs = nil
			} else {
				if z.Airs == nil {
					z.Airs = new(Airs)
				}
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Day":
						z.Airs.Day, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Time":
						z.Airs.Time, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Timezone":
						z.Airs.Timezone, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Genres":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0003) {
				z.Genres = (z.Genres)[:zb0003]
			} else {
				z.Genres = make([]string, zb0003)
			}
			for za0001 := range z.Genres {
				z.Genres[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Country":
			z.Country, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Language":
			z.Language, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Translations":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Translations) >= int(zb0004) {
				z.Translations = (z.Translations)[:zb0004]
			} else {
				z.Translations = make([]string, zb0004)
			}
			for za0002 := range z.Translations {
				z.Translations[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Show) Msgsize() (s int) {
	s = 3 + 7 + z.Object.Msgsize() + 11 + msgp.StringPrefixSize + len(z.FirstAired) + 4 + msgp.StringPrefixSize + len(z.URL) + 8 + msgp.StringPrefixSize + len(z.Trailer) + 8 + msgp.IntSize + 9 + msgp.StringPrefixSize + len(z.Overview) + 14 + msgp.StringPrefixSize + len(z.Certification) + 7 + msgp.StringPrefixSize + len(z.Status) + 8 + msgp.StringPrefixSize + len(z.Network) + 14 + msgp.IntSize + 5
	if z.Airs == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 4 + msgp.StringPrefixSize + len(z.Airs.Day) + 5 + msgp.StringPrefixSize + len(z.Airs.Time) + 9 + msgp.StringPrefixSize + len(z.Airs.Timezone)
	}
	s += 7 + msgp.Float32Size + 6 + msgp.IntSize + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Genres {
		s += msgp.StringPrefixSize + len(z.Genres[za0001])
	}
	s += 8 + msgp.StringPrefixSize + len(z.Country) + 9 + msgp.StringPrefixSize + len(z.Language) + 13 + msgp.ArrayHeaderSize
	for za0002 := range z.Translations {
		s += msgp.StringPrefixSize + len(z.Translations[za0002])
	}
	s += 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ShowSearchResults) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 3
		// string "Type"
		o = append(o, 0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
		o = msgp.AppendString(o, z[za0001].Type)
		// string "Score"
		o = append(o, 0xa5, 0x53, 0x63, 0x6f, 0x72, 0x65)
		o, err = msgp.AppendIntf(o, z[za0001].Score)
		if err != nil {
			return
		}
		// string "Show"
		o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
		if z[za0001].Show == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].Show.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ShowSearchResults) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ShowSearchResults, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch msgp.UnsafeString(field) {
			case "Type":
				(*z)[zb0001].Type, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "Score":
				(*z)[zb0001].Score, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			case "Show":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].Show = nil
				} else {
					if (*z)[zb0001].Show == nil {
						(*z)[zb0001].Show = new(Show)
					}
					bts, err = (*z)[zb0001].Show.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ShowSearchResults) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + 5 + msgp.StringPrefixSize + len(z[zb0004].Type) + 6 + msgp.GuessSize(z[zb0004].Score) + 5
		if z[zb0004].Show == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].Show.Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Shows) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Watchers"
	o = append(o, 0x82, 0xa8, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	o = msgp.AppendInt(o, z.Watchers)
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Shows) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Watchers":
			z.Watchers, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Shows) Msgsize() (s int) {
	s = 1 + 9 + msgp.IntSize + 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Sizes) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Full"
	o = append(o, 0x83, 0xa4, 0x46, 0x75, 0x6c, 0x6c)
	o = msgp.AppendString(o, z.Full)
	// string "Medium"
	o = append(o, 0xa6, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d)
	o = msgp.AppendString(o, z.Medium)
	// string "Thumbnail"
	o = append(o, 0xa9, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.Thumbnail)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sizes) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Full":
			z.Full, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Medium":
			z.Medium, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Thumbnail":
			z.Thumbnail, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sizes) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Full) + 7 + msgp.StringPrefixSize + len(z.Medium) + 10 + msgp.StringPrefixSize + len(z.Thumbnail)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Token) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "AccessToken"
	o = append(o, 0x85, 0xab, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.AccessToken)
	// string "TokenType"
	o = append(o, 0xa9, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.TokenType)
	// string "ExpiresIn"
	o = append(o, 0xa9, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e)
	o = msgp.AppendInt(o, z.ExpiresIn)
	// string "RefreshToken"
	o = append(o, 0xac, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.RefreshToken)
	// string "Scope"
	o = append(o, 0xa5, 0x53, 0x63, 0x6f, 0x70, 0x65)
	o = msgp.AppendString(o, z.Scope)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Token) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AccessToken":
			z.AccessToken, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TokenType":
			z.TokenType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ExpiresIn":
			z.ExpiresIn, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "RefreshToken":
			z.RefreshToken, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Scope":
			z.Scope, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Token) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.AccessToken) + 10 + msgp.StringPrefixSize + len(z.TokenType) + 10 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.RefreshToken) + 6 + msgp.StringPrefixSize + len(z.Scope)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TokenRefresh) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "RefreshToken"
	o = append(o, 0x85, 0xac, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.RefreshToken)
	// string "ClientID"
	o = append(o, 0xa8, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44)
	o = msgp.AppendString(o, z.ClientID)
	// string "ClientSecret"
	o = append(o, 0xac, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.ClientSecret)
	// string "RedirectURI"
	o = append(o, 0xab, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x52, 0x49)
	o = msgp.AppendString(o, z.RedirectURI)
	// string "GrantType"
	o = append(o, 0xa9, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.GrantType)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TokenRefresh) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RefreshToken":
			z.RefreshToken, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ClientID":
			z.ClientID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ClientSecret":
			z.ClientSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RedirectURI":
			z.RedirectURI, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "GrantType":
			z.GrantType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TokenRefresh) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.RefreshToken) + 9 + msgp.StringPrefixSize + len(z.ClientID) + 13 + msgp.StringPrefixSize + len(z.ClientSecret) + 12 + msgp.StringPrefixSize + len(z.RedirectURI) + 10 + msgp.StringPrefixSize + len(z.GrantType)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserSettings) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "User"
	// map header, size 2
	// string "Username"
	o = append(o, 0x82, 0xa4, 0x55, 0x73, 0x65, 0x72, 0x82, 0xa8, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.User.Username)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.User.Name)
	// string "Account"
	// map header, size 0
	o = append(o, 0xa7, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserSettings) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "User":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Username":
					z.User.Username, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Name":
					z.User.Name, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "Account":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserSettings) Msgsize() (s int) {
	s = 1 + 5 + 1 + 9 + msgp.StringPrefixSize + len(z.User.Username) + 5 + msgp.StringPrefixSize + len(z.User.Name) + 8 + 1
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchedItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "MediaType"
	o = append(o, 0x87, 0xa9, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.MediaType)
	// string "KodiID"
	o = append(o, 0xa6, 0x4b, 0x6f, 0x64, 0x69, 0x49, 0x44)
	o = msgp.AppendInt(o, z.KodiID)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	o = msgp.AppendInt(o, z.Movie)
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	o = msgp.AppendInt(o, z.Show)
	// string "Season"
	o = append(o, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	o = msgp.AppendInt(o, z.Season)
	// string "Episode"
	o = append(o, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, z.Episode)
	// string "Watched"
	o = append(o, 0xa7, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Watched)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchedItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "MediaType":
			z.MediaType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "KodiID":
			z.KodiID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			z.Movie, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Show":
			z.Show, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Season":
			z.Season, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Episode":
			z.Episode, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Watched":
			z.Watched, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchedItem) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.MediaType) + 7 + msgp.IntSize + 6 + msgp.IntSize + 5 + msgp.IntSize + 7 + msgp.IntSize + 8 + msgp.IntSize + 8 + msgp.BoolSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchedMovie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Plays"
	o = append(o, 0x83, 0xa5, 0x50, 0x6c, 0x61, 0x79, 0x73)
	o = msgp.AppendInt(o, z.Plays)
	// string "LastWatchedAt"
	o = append(o, 0xad, 0x4c, 0x61, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendTime(o, z.LastWatchedAt)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchedMovie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Plays":
			z.Plays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "LastWatchedAt":
			z.LastWatchedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchedMovie) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 14 + msgp.TimeSize + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchedProgressShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Aired"
	o = append(o, 0x87, 0xa5, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendInt(o, z.Aired)
	// string "Completed"
	o = append(o, 0xa9, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64)
	o = msgp.AppendInt(o, z.Completed)
	// string "LastWatchedAt"
	o = append(o, 0xad, 0x4c, 0x61, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendTime(o, z.LastWatchedAt)
	// string "Seasons"
	o = append(o, 0xa7, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons)))
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Seasons[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "HiddenSeasons"
	o = append(o, 0xad, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.HiddenSeasons)))
	for za0002 := range z.HiddenSeasons {
		if z.HiddenSeasons[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.HiddenSeasons[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "NextEpisode"
	o = append(o, 0xab, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.NextEpisode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.NextEpisode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "LastEpisode"
	o = append(o, 0xab, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.LastEpisode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.LastEpisode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchedProgressShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Aired":
			z.Aired, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Completed":
			z.Completed, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "LastWatchedAt":
			z.LastWatchedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Seasons":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seasons) >= int(zb0002) {
				z.Seasons = (z.Seasons)[:zb0002]
			} else {
				z.Seasons = make([]*Season, zb0002)
			}
			for za0001 := range z.Seasons {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Seasons[za0001] = nil
				} else {
					if z.Seasons[za0001] == nil {
						z.Seasons[za0001] = new(Season)
					}
					bts, err = z.Seasons[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "HiddenSeasons":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.HiddenSeasons) >= int(zb0003) {
				z.HiddenSeasons = (z.HiddenSeasons)[:zb0003]
			} else {
				z.HiddenSeasons = make([]*Season, zb0003)
			}
			for za0002 := range z.HiddenSeasons {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.HiddenSeasons[za0002] = nil
				} else {
					if z.HiddenSeasons[za0002] == nil {
						z.HiddenSeasons[za0002] = new(Season)
					}
					bts, err = z.HiddenSeasons[za0002].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "NextEpisode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.NextEpisode = nil
			} else {
				if z.NextEpisode == nil {
					z.NextEpisode = new(Episode)
				}
				bts, err = z.NextEpisode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "LastEpisode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LastEpisode = nil
			} else {
				if z.LastEpisode == nil {
					z.LastEpisode = new(Episode)
				}
				bts, err = z.LastEpisode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchedProgressShow) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 10 + msgp.IntSize + 14 + msgp.TimeSize + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Seasons[za0001].Msgsize()
		}
	}
	s += 14 + msgp.ArrayHeaderSize
	for za0002 := range z.HiddenSeasons {
		if z.HiddenSeasons[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.HiddenSeasons[za0002].Msgsize()
		}
	}
	s += 12
	if z.NextEpisode == nil {
		s += msgp.NilSize
	} else {
		s += z.NextEpisode.Msgsize()
	}
	s += 12
	if z.LastEpisode == nil {
		s += msgp.NilSize
	} else {
		s += z.LastEpisode.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchedShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Plays"
	o = append(o, 0x85, 0xa5, 0x50, 0x6c, 0x61, 0x79, 0x73)
	o = msgp.AppendInt(o, z.Plays)
	// string "Watched"
	o = append(o, 0xa7, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Watched)
	// string "LastWatchedAt"
	o = append(o, 0xad, 0x4c, 0x61, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendTime(o, z.LastWatchedAt)
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Seasons"
	o = append(o, 0xa7, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons)))
	for za0001 := range z.Seasons {
		// map header, size 3
		// string "Plays"
		o = append(o, 0x83, 0xa5, 0x50, 0x6c, 0x61, 0x79, 0x73)
		o = msgp.AppendInt(o, z.Seasons[za0001].Plays)
		// string "Number"
		o = append(o, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
		o = msgp.AppendInt(o, z.Seasons[za0001].Number)
		// string "Episodes"
		o = append(o, 0xa8, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons[za0001].Episodes)))
		for za0002 := range z.Seasons[za0001].Episodes {
			// map header, size 3
			// string "Number"
			o = append(o, 0x83, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
			o = msgp.AppendInt(o, z.Seasons[za0001].Episodes[za0002].Number)
			// string "Plays"
			o = append(o, 0xa5, 0x50, 0x6c, 0x61, 0x79, 0x73)
			o = msgp.AppendInt(o, z.Seasons[za0001].Episodes[za0002].Plays)
			// string "LastWatchedAt"
			o = append(o, 0xad, 0x4c, 0x61, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74)
			o = msgp.AppendTime(o, z.Seasons[za0001].Episodes[za0002].LastWatchedAt)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchedShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Plays":
			z.Plays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Watched":
			z.Watched, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "LastWatchedAt":
			z.LastWatchedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Seasons":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seasons) >= int(zb0002) {
				z.Seasons = (z.Seasons)[:zb0002]
			} else {
				z.Seasons = make([]struct {
					Plays    int `json:"plays"`
					Number   int `json:"number"`
					Episodes []struct {
						Number        int       `json:"number"`
						Plays         int       `json:"plays"`
						LastWatchedAt time.Time `json:"last_watched_at"`
					} `json:"episodes"`
				}, zb0002)
			}
			for za0001 := range z.Seasons {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Plays":
						z.Seasons[za0001].Plays, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					case "Number":
						z.Seasons[za0001].Number, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					case "Episodes":
						var zb0004 uint32
						zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Seasons[za0001].Episodes) >= int(zb0004) {
							z.Seasons[za0001].Episodes = (z.Seasons[za0001].Episodes)[:zb0004]
						} else {
							z.Seasons[za0001].Episodes = make([]struct {
								Number        int       `json:"number"`
								Plays         int       `json:"plays"`
								LastWatchedAt time.Time `json:"last_watched_at"`
							}, zb0004)
						}
						for za0002 := range z.Seasons[za0001].Episodes {
							var zb0005 uint32
							zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
							for zb0005 > 0 {
								zb0005--
								field, bts, err = msgp.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								switch msgp.UnsafeString(field) {
								case "Number":
									z.Seasons[za0001].Episodes[za0002].Number, bts, err = msgp.ReadIntBytes(bts)
									if err != nil {
										return
									}
								case "Plays":
									z.Seasons[za0001].Episodes[za0002].Plays, bts, err = msgp.ReadIntBytes(bts)
									if err != nil {
										return
									}
								case "LastWatchedAt":
									z.Seasons[za0001].Episodes[za0002].LastWatchedAt, bts, err = msgp.ReadTimeBytes(bts)
									if err != nil {
										return
									}
								default:
									bts, err = msgp.Skip(bts)
									if err != nil {
										return
									}
								}
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchedShow) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 8 + msgp.BoolSize + 14 + msgp.TimeSize + 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Seasons {
		s += 1 + 6 + msgp.IntSize + 7 + msgp.IntSize + 9 + msgp.ArrayHeaderSize + (len(z.Seasons[za0001].Episodes) * (28 + msgp.IntSize + msgp.IntSize + msgp.TimeSize))
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Watchlist) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Movies"
	o = append(o, 0x83, 0xa6, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Movies)))
	for za0001 := range z.Movies {
		if z.Movies[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Movies[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Shows"
	o = append(o, 0xa5, 0x53, 0x68, 0x6f, 0x77, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Shows)))
	for za0002 := range z.Shows {
		if z.Shows[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Shows[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Episodes"
	o = append(o, 0xa8, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Episodes)))
	for za0003 := range z.Episodes {
		if z.Episodes[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Episodes[za0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Watchlist) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Movies":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Movies) >= int(zb0002) {
				z.Movies = (z.Movies)[:zb0002]
			} else {
				z.Movies = make([]*Movie, zb0002)
			}
			for za0001 := range z.Movies {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Movies[za0001] = nil
				} else {
					if z.Movies[za0001] == nil {
						z.Movies[za0001] = new(Movie)
					}
					bts, err = z.Movies[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Shows":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Shows) >= int(zb0003) {
				z.Shows = (z.Shows)[:zb0003]
			} else {
				z.Shows = make([]*Show, zb0003)
			}
			for za0002 := range z.Shows {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Shows[za0002] = nil
				} else {
					if z.Shows[za0002] == nil {
						z.Shows[za0002] = new(Show)
					}
					bts, err = z.Shows[za0002].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Episodes":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Episodes) >= int(zb0004) {
				z.Episodes = (z.Episodes)[:zb0004]
			} else {
				z.Episodes = make([]*Episode, zb0004)
			}
			for za0003 := range z.Episodes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Episodes[za0003] = nil
				} else {
					if z.Episodes[za0003] == nil {
						z.Episodes[za0003] = new(Episode)
					}
					bts, err = z.Episodes[za0003].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Watchlist) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Movies {
		if z.Movies[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Movies[za0001].Msgsize()
		}
	}
	s += 6 + msgp.ArrayHeaderSize
	for za0002 := range z.Shows {
		if z.Shows[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.Shows[za0002].Msgsize()
		}
	}
	s += 9 + msgp.ArrayHeaderSize
	for za0003 := range z.Episodes {
		if z.Episodes[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += z.Episodes[za0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchlistEpisode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ListedAt"
	o = append(o, 0x84, 0xa8, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.ListedAt)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Episode"
	o = append(o, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.Episode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Episode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchlistEpisode) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ListedAt":
			z.ListedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Episode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Episode = nil
			} else {
				if z.Episode == nil {
					z.Episode = new(Episode)
				}
				bts, err = z.Episode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Object)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchlistEpisode) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.ListedAt) + 5 + msgp.StringPrefixSize + len(z.Type) + 8
	if z.Episode == nil {
		s += msgp.NilSize
	} else {
		s += z.Episode.Msgsize()
	}
	s += 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchlistMovie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "ListedAt"
	o = append(o, 0x83, 0xa8, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.ListedAt)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Movie"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchlistMovie) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ListedAt":
			z.ListedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(Movie)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchlistMovie) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.ListedAt) + 5 + msgp.StringPrefixSize + len(z.Type) + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchlistSeason) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ListedAt"
	o = append(o, 0x84, 0xa8, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.ListedAt)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Season"
	o = append(o, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	if z.Season == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Season.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchlistSeason) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ListedAt":
			z.ListedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Season":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Season = nil
			} else {
				if z.Season == nil {
					z.Season = new(Object)
				}
				bts, err = z.Season.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Object)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchlistSeason) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.ListedAt) + 5 + msgp.StringPrefixSize + len(z.Type) + 7
	if z.Season == nil {
		s += msgp.NilSize
	} else {
		s += z.Season.Msgsize()
	}
	s += 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WatchlistShow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "ListedAt"
	o = append(o, 0x83, 0xa8, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74)
	o = msgp.AppendString(o, z.ListedAt)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Show"
	o = append(o, 0xa4, 0x53, 0x68, 0x6f, 0x77)
	if z.Show == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Show.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WatchlistShow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ListedAt":
			z.ListedAt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Show":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Show = nil
			} else {
				if z.Show == nil {
					z.Show = new(Show)
				}
				bts, err = z.Show.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WatchlistShow) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.ListedAt) + 5 + msgp.StringPrefixSize + len(z.Type) + 5
	if z.Show == nil {
		s += msgp.NilSize
	} else {
		s += z.Show.Msgsize()
	}
	return
}
