package chardet

var dict_b5 = map[uint32]int{
	0xAABA: 0x00, // 的
	0xA440: 0x01, // 一
	0xA662: 0x02, // 在
	0xA448: 0x03, // 人
	0xA446: 0x04, // 了
	0xA6B3: 0x05, // 有
	0xA4A4: 0x06, // 中
	0xAC4F: 0x07, // 是
	0xA67E: 0x08, // 年
	0xA94D: 0x09, // 和
	0xA46A: 0x0A, // 大
	0xA4A3: 0x0B, // 不
	0xA475: 0x0C, // 工
	0xA457: 0x0D, // 上
	0xA661: 0x0E, // 地
	0xA5AB: 0x0F, // 市
	0xAD6E: 0x10, // 要
	0xA558: 0x11, // 出
	0xA6E6: 0x12, // 行
	0xA740: 0x13, // 作
	0xA5CD: 0x14, // 生
	0xAE61: 0x15, // 家
	0xA548: 0x16, // 以
	0xA6A8: 0x17, // 成
	0xA8EC: 0x18, // 到
	0xA4E9: 0x19, // 日
	0xA5C1: 0x1A, // 民
	0xA7DA: 0x1B, // 我
	0xB3A1: 0x1C, // 部
	0xA668: 0x1D, // 多
	0xA5FE: 0x1E, // 全
	0xABD8: 0x1F, // 建
	0xA54C: 0x20, // 他
	0xA4BD: 0x21, // 公
	0xAE69: 0x22, // 展
	0xB27A: 0x23, // 理
	0xB773: 0x24, // 新
	0xA4E8: 0x25, // 方
	0xA544: 0x26, // 主
	0xA5F8: 0x27, // 企
	0xA8EE: 0x28, // 制
	0xAC46: 0x29, // 政
	0xA5CE: 0x2A, // 用
	0xA650: 0x2B, // 同
	0xAA6B: 0x2C, // 法
	0xB0AA: 0x2D, // 高
	0xA5BB: 0x2E, // 本
	0xA4EB: 0x2F, // 月
	0xA977: 0x30, // 定
	0xA4C6: 0x31, // 化
	0xA55B: 0x32, // 加
	0xA658: 0x33, // 合
	0xAB7E: 0x34, // 品
	0xADAB: 0x35, // 重
	0xA4C0: 0x36, // 分
	0xA44F: 0x37, // 力
	0xA57E: 0x38, // 外
	0xB44E: 0x39, // 就
	0xB5A5: 0x3A, // 等
	0xA455: 0x3B, // 下
	0xA4B8: 0x3C, // 元
	0xAAC0: 0x3D, // 社
	0xAB65: 0x3E, // 前
	0xADB1: 0x3F, // 面
	0xA45D: 0x40, // 也
	0xA4A7: 0x41, // 之
	0xA6D3: 0x42, // 而
	0xA751: 0x43, // 利
	0xA4E5: 0x44, // 文
	0xA8C6: 0x45, // 事
	0xA569: 0x46, // 可
	0xA7EF: 0x47, // 改
	0xA655: 0x48, // 各
	0xA66E: 0x49, // 好
	0xAAF7: 0x4A, // 金
	0xA571: 0x4B, // 司
	0xA8E4: 0x4C, // 其
	0xA5AD: 0x4D, // 平
	0xA54E: 0x4E, // 代
	0xA4D1: 0x4F, // 天
}

// [\x00-\x7F]
// [\xA1-\xF9][\x40-\x7E\xA1-\xFE]
type big5 struct {
	byte
	rune
	hold [80]int
	ttls int
}

func (b big5) String() string {
	return "big5"
}

func (b *big5) Feed(x byte) (ans bool) {
	if b.byte == 0 {
		if x <= 0x7F {
			return true
		}
		if x >= 0xA1 && x <= 0xF9 {
			b.byte = 1
			b.rune = rune(x) << 8
			return true
		}
	} else {
		if (x >= 0x40 && x <= 0x7E) || (x >= 0xA1 && x <= 0xFE) {
			b.byte = 0
			b.rune |= rune(x)
			b.count()
			return true
		}
	}
	return false
}

func (b *big5) Priority() float64 {
	if b.ttls == 0 {
		return 0
	}
	f := 0.0
	for i, x := range b.hold {
		k := 100*float64(x)/float64(b.ttls) - freq_ch[i]
		if k >= 0 {
			f += k
		} else {
			f -= k
		}
	}
	return 100 - f
}

func (b *big5) count() {
	if i, ok := dict_b5[uint32(b.rune)]; ok {
		b.hold[i]++
		b.ttls++
	}
}
