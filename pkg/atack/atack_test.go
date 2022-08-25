package atack

import "testing"

func TestAtack_buildString(t *testing.T) {
	type fields struct {
		atack     Atacker
		pass      string
		maxLength int
		lenght    int
		chars     []rune
	}

	chars := []rune{'a', 'b', 'c'}

	type args struct {
		i []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "aaa test",
			fields: fields{
				atack:     nil,
				pass:      "",
				maxLength: 3,
				lenght:    3,
				chars:     chars,
			},
			args: args{[]int{0, 0, 0}},
			want: "aaa",
		},
		{
			name: "abc test",
			fields: fields{
				atack:     nil,
				pass:      "",
				maxLength: 3,
				lenght:    3,
				chars:     chars,
			},
			args: args{[]int{0, 1, 2}},
			want: "abc",
		}, {
			name: "abcabcabc test",
			fields: fields{
				atack:     nil,
				pass:      "",
				maxLength: 9,
				lenght:    9,
				chars:     chars,
			},
			args: args{[]int{0, 1, 2, 0, 1, 2, 0, 1, 2}},
			want: "abcabcabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Atack{
				atack:     tt.fields.atack,
				maxLength: tt.fields.maxLength,
				lenght:    tt.fields.lenght,
				chars:     tt.fields.chars,
			}
			ans, _ := a.buildString(tt.args.i)
			if got := ans; got != tt.want {
				t.Errorf("buildString() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockAtacker struct {
	secret string
}

func newMockAtacker(secret string) *mockAtacker {
	return &mockAtacker{secret: secret}
}
func (m *mockAtacker) prepare() error {
	return nil
}

func (m *mockAtacker) check(pass string) bool {
	return m.secret == pass
}

func (m *mockAtacker) setSecret(pass string) {
	m.secret = pass
}

func TestAtack_brute(t *testing.T) {

	ma := newMockAtacker("abc")
	mb := newMockAtacker("aaa")
	mc := newMockAtacker("ccc")
	me := newMockAtacker("ccc")

	type fields struct {
		atack     Atacker
		maxLength int
		chars     []rune
	}
	tests := []struct {
		name     string
		fields   fields
		wantPass string
		wantErr  bool
	}{
		{
			name: "Pass abc",
			fields: fields{
				atack:     ma,
				maxLength: 3,
				chars:     []rune{'a', 'b', 'c'},
			},
			wantPass: "abc",
			wantErr:  false,
		},
		{
			name: "Pass aaa",
			fields: fields{
				atack:     mb,
				maxLength: 3,
				chars:     []rune{'a', 'b', 'c'},
			},
			wantPass: "aaa",
			wantErr:  false,
		},
		{
			name: "Pass ccc",
			fields: fields{
				atack:     mc,
				maxLength: 3,
				chars:     []rune{'a', 'b', 'c'},
			},
			wantPass: "ccc",
			wantErr:  false,
		},
		{
			name: "Pass ccc in 4 chars pass",
			fields: fields{
				atack:     me,
				maxLength: 4,
				chars:     []rune{'c', 'c', 'c'},
			},
			wantPass: "ccc",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAtack(tt.fields.atack,
				tt.fields.maxLength,
				tt.fields.chars)

			gotPass, err := a.Brute()
			if (err != nil) != tt.wantErr {
				t.Errorf("brute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPass != tt.wantPass {
				t.Errorf("brute() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
