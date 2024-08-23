package main

import "testing"

func Test_Soma_Correct(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Soma_Incorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Mult_Correct(t *testing.T) {
	teste := multiplica(3, 3)
	resultado := 9
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Mult_Incorrect(t *testing.T) {
	teste := multiplica(3, 3)
	resultado := 11
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Diminuicao_Correct(t *testing.T) {
	teste := diminui(100, 20)
	resultado := 80
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Diminuicao_Incorrect(t *testing.T) {
	teste := diminui(100, 20)
	resultado := 30
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Div_Correct(t *testing.T) {
	teste := divide(100, 25)
	resultado := 4
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}

func Test_Div_Incorrect(t *testing.T) {
	teste := divide(100, 25)
	resultado := 5
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor obtido: %d", resultado, teste)
	}
}
