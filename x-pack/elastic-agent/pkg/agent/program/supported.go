// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("eJzUelt3oziX9v33M/r2m5mXQ0g3s1Zf2E6DwQ4p40QSukPCBmyB6RgfYNb891kSZ+xUVapq5p25qFUJETps7f3sZz+b//jtmG3oP/ws+dfj5v28ef+3ImG//ftvJDFy/HoIV2DqLIHDaIoZDbMdgatHyzQuZC2XGNkKRtbCQ7bkQxx56t2/pbQ8hPByCK2Zlbtr62jN7NyDWoQVkGOoScsEnDxoHzFc6cHclvHaOs7iSWjFsmHFl9BK5IgkRoqhzEi6OhGo7/Gr/LBE08hTjjlOrgyjVbycDNaIqCnlVAFFkIBimbhRYBpbqroFhkZOC42RxIiJCfZfkFsE8G38furJ+gUje8f3+GXt/s5t8SaBNdiDNZKCL670/Pj8Orls19YCrbIigFdGy8PCmk1yy3TOdM62galvicnK4Ek8D/k/H7mMvB7Cet/t88D849Ey2YnOgUSFvfjz6ZkwfUcULSEm2wVPh5Ak+pmKvUohgcblJWzHXjzkHl7iaUJMwIJJ+7y/t/rZJLfmNiNQVzDQ3zHaPwqbz3r/zJxtXg+hB50dRk6JFOPiV2N/aB6agAtRua9oaX/Mcj0tMJTPQQK2vgK07jzNv6mY1zK5D2gSf3e5niY+itr7fImnv29XlU1wwo6BCQqk4ojMweDcVAFHDB2JqNY9W1frzF2Nmm/dvSTGMYCgHO5LCj00zdu9wKuEkdWdfe4wYoJdYOrFjd1v1q3mIyqQ+PORbeq7zFTfBKeXeHrEUEsDMzzY8zwj6VQO5o6+WE/+v/U0CT2o7S0ziqiUs8063G8UXQ7mfIx0tGYBI6ZRBibbUQVENHEOdrFf/PYvFQhs0iA7xGk+ggAXantq6hlJV+GbAnYBsrNgvl94irx/iaeMJO6FKOwUzOQSQ0emCZM2qyyiqZvhxBAui7s5cmwCZZYKSMk85e3RevLUl6dw4UFH8qF+QkplHqTyaxBmzy0TnPB8evb59SfXM5Z14eqVaad7D9mqDx8erZl1fjVZTBOj2Kx1oznuUureX6qO5CGXLZXrGRd6b//S30s+d2HxOY8+1OTN0yG0Yv1M56uzC68RVd3MK3Sje0cvA9OQ8Fo/EoWe++dcxBp/FvPrDBR2wqaucii09s+PyLiuaKKnNDFy6y+cEROUyLi2+xU/N2sYV0pVDl+AIpOf/UrvrpM4Bwydd2E/1Y2IeXmcxVKIUcQ8WecuyhoXpqYh+fxsSc8uyGGeCgofuZpVj6vhe9G4qMXTQMKSzdrqnsVSTkw9bd5ZricxVd09RnbRPAtMlmOoy9wXnsvJgpp6GRh8/47kweuxvuMHDJ0tD1/8Wq85n0aBGT5aM/u+nzX7MI0Cq89N+OXWzG7n7u9ruZbbO6nHlYHpMppavWdWvkTgglU7wubb6LnNqKLLNHEYLXo2+MCOw/Hao48m9XxTyefpTAXSSzxRnp8mCzq3GVLByYca96kjeToslusp25hghxTuI2/1+RqYn8R9P6BdbDZrRDQJ+rDFzyuTpPWPuIOa23u8b587+07BEaPnytcS9oBUR6IJiHiK26jSon5+H5J7MPkST2s4e17000YwZxe8OtyFYm6f1i8mwl41FNvb8divQb9IOwLLKkpQn1Xx4FWuztakOkEVxJn6qbeJp0FKn9syMQd7/TjNNSmypiszEQ84I3OX0d2trfoxOaQRUuhD7RIgt2z3bNYpYNLbB8IZVdiZhIdFoESM7A4h4RiruofFzP29mtMdpZUrI0kg+TOeVmr7qVJmPT2Ez7NpRJJV6JtGuVaAxudo0tl2fQltBRw9xPHdKTE0Ck8J08Uq47SmDOZ2xOOGY6M9z4sAasLHlgkfF+nW7C/dmgURTSTFmdE2XW1jtiEb/yZdcfiANvPQqklRAvq8BETBJKtCIp6SAcNMHRbMwWWZsCNZ95gh5O7qMCu+cPPGy7e3eDmbxFQBUoAmp8AEOTWvUWC+nTDUIo+b7UlOPHgtfxGL/SRTvp6xEmQkoSdiAgXDi45NEAeQfpPdunv27L7Zxgq4W1fS56vycH1+mupWmAWzxCgCkyU8Rb/E0xhDQ6KF3rdrgRFPq7aEFJ7Ca5cz9J1vGifM03wdOt+Ci+V6mpEkY57qbn2o7TEKG9iTNmjKOAMiSlBWIcJOfgJ2gaFzN2IelLd0bp89BZRU0YsmxIiibT1FP+HkmlVVCjuJUDP0CKcuo02a4lDzykNBT+lFwEnuockIdkcwdRvmO6JONaQYR2LoEpH1o48caRDuc/f8Ek+bPZd9xne7V+1M1EmPQfKU4DA6Xz1axjHuKhpxrxmeif9bH678e1oQxWFUdc40fY7RarRX1T0j5ZpRdTVg702l0rfpgLV+6hytzWMM8ZkmFYxS4dNyS01wap/5mkMYm0qk7J61Z77UKXLuFtw3qzO4W7HvFrbH9+aM9jvNeBzfnmO4ZuVvckSfRtRklEo6/+5Shge1knCa+NqvBNp9Cb8e2A4aF2qCYlwJcb+s6ZH4eVhpTVmQAE5JK6on4oQuhpUYx4SKAmOTncSZTTcjo3UwwhyjzoHpXPj+vIH/9ecBe/p0CAPoXkZ7qeiwqe98BfDUvieK846RNZpHUPACQzejsl4SU1f5uV7iafXscnv+pepolFP38hBSlZU9O3y1Eq1TvEwVsB1VWt/z3t5DbtRi1er7Uzo2DckDDT5JJ5GnkHPxoMN+eP21Jn7neWhAFZg+oizTEnOMVt0z3d3E1G21Ox9XzLzKDcoldI5EBftAMSRPCYd/Q889qmE8BshmS1j5hz8HHX02qzlwosskcYtNH4MUTfIgO3X+IYWN73TvBzs6ky8BtMsAduP8uStRQ+/TwGbeQzB3L37qnEmHze8exO9et3YvRqSQIEPqvX/CyE24T3Z7+KN0diAKoJsN6FTPV5frYUzx3zfIGazD46q7e7Cn3Vycnj90f+O5/HrG3dwithv8++m7N1s+8eH9C56xOgzysgcDRlKRnxOOkQ3N9xT9smkwzsQFUaQbfKxULlHyCh7xOarbcbsPVJNcxB3iNN9mszQ4YPjwOKSu3drL5Odp7HI2SWkC9j56TiuOFVQ+tqZHaxYIrhKYRunPaDYL//yzpa9ss8nvi65uRQXDt6YaSJwcd9VFK2pyWsarXk4DlmuZ00ReHedEcZnFpJEwi48YAlGd1mLuSLDE2aalUdKjNc/1Wfg/K6jdnO9XrT9286/soXH3r8JyDe2t69X7bPaCOAyZf3wkBoawqEW9mdYI5admrmVy42JhQ9d6lOEcIPcSoNWiFiCbUOchyWnLHfvYouxp/SCt/SDWLkThFG1/8uHq3loN3T09z9qxn1v3/j2fOnE+OlN1NRJJ7+xTte+Loz07UtU5dnYfUbTi4dLOFzZ+36N0H1Xyt7SuD8dD4Xf2TQG4d8ZO0B7/nacvojqtb30g9A9LkY/3+X0C+xyUGOitb1k/IPKP6M6vmOMeZfrcucxh86duBpX4B8/YLz8a7Bykr+/z/881BcK7DYDK/+fD9NVXcpq0v13vwy/x5GLxcnw2PXjIWWK053PU9+3qPJXhRuVWHeYhezdOW9HGf8/vyC5rE0Q0dauUXucrf/Csl6tGMooPr3lf2sCJcaRKNeazksvH0ol8xuabvpnJ3I/el7DCF1EefyDR9OaXaArYcOzHUgsadT+ERDJ/Pg/skdqcxvHft7xUbu4ZKc7Bg1qKhfLJc9F35OlBE6pTPpEaZIEZbWkCUoyiPn7czcs0eXu0DK3ccOpuApGDlsXD+708NsKer9DQb6i8H6qv9+hoE6eVlPQjsTdSRxscuY9536SS/zz6mGzy95jeCcRXCCSasF2ty9Ude5kFczvzlFoXvd+VLzFyZTrTMmJK39YzG800dRlB06Noqd0L4q/qmR8G5UBvJdA4iZoDcoL7zc49N1rqQy1dJlcWJOD4BbrMS0E6nhcrzpaaoPA4cX6S9htJNoBh2yvpUzpojpFb+ND5bGBfeIIftuuEM+ZExWzY8hm0wjjBF5rjMlmdqcpKDorLlOVkZvW+JjAKrIg2TUygvhcgUmh7oZ+oLiNrTSPwcsImK+s7i2dpq2V+ADQigP5uSC42wcPLWDut/a/W5tqasUmaSO1pkU3QmvoJKc6ZJPjoV7bNiCBqvNbEUqsLVIDV+bP6/Ks0xK/ovDd64ehrirruBXpEzOvdLz6atXtr9kjFzdlPRNEHehpG0Q6jqSQIfFprQOhZ+LAPV5UvN/Fa6b8Xmujiqw0OgrwYHO31vi5tMgmLux7oxT96ju4OE5AIEK3bvMHcjrhPCDsZutBfu5bgiLSr7XkXwz2OnldadJtc7rYdZV31kXtAqs2wAh5a4n+n1Vglw4fToovv0p58tZX4T28/jvz9xk+53UnqCKz+mBz07mBEEPicPMZRHP7j7ekqdOsv8cP7Yn1ro2oevgbHLLevcVdEu8pX/bkTmuj5PT280ozkM56DI+60u9yDOUOKUdDE0O76cYsTTu9rqNZX2j3jfrv021ps771Oh6IJqHFb+NLitggZixDf886wcPklevGPzyEwlqrTyFPevv8Mo6JymEOqAljkkuSPRnSqv55r9tT6xajvNrjDTksc+1arQw440Olebmjz5CCWrVs9Ne3F5Ihb3RRu93TJ+8Xc4CyCG3F/Rz/zdZfgOx3BnE1SXzESX/lL/Cz6nJyriHb8oSWXh+Pfp817cY9dqs41gKDYDDvsZ6oaMka2Nu6yf6LD/nlm+TFDFBLoFxTm1DR2fiHXMtHXysPbzvp3MUomLWgCOEUvAkM/E9Z4urv1lCgiScA4ovHy8QP2eGM7T3UzzhxrD92KrFR3zElj/9Xh/1on/WcyYy8j3MmKpl5SBBhN94Osfzu/3e/Sj8rQb0otv+RbS247juYeZKdg/swz4+80BSfxgdxaE+O260n0ZT1JPWTnzm6VVpLNZL9RbZmkbkbg20lIM3MptF6l0C6fH5y4+zAm8+l+c0+jeePBoABpUBrOOZznLDBHpWFBc7eqQb9RFvIxN2MlDOWLaHnfCXa7uIR2IRvif+W7PosRcLV8e6tg6+vl4HCsqaeYp5dCO4q29pO8x9CWcRF8VM6JUhSjn9dpflIPGdKgD7WQ+v4mzbptIFf9IEPPSILP9QfS/y2ayf8ObeTw52//+f/+KwAA///73WjS")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
