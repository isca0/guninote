### guninote

Efetua o cálculo de média da faculdade uninter.

Este script em go, efetua de forma simples o cálculo de média de suas matérias da  
da faculdade uninter.  

### Como Usar:
  
`go get github.com/isca0/guninter`
  
`guninter`

O script aguarda que você entre com suas notas de Apols, prova prática, prova discursiva e  
ativida prática.  

Apenas entre os dados conforme for lhe solicitado e pressione enter, ao final você receberá o  
resultado de sua média. :wink:  
  
### Cálculo

O cálculo de média segue esta lógica aritimética:
  
_Para obter o resultado da média das apols o cálculo é feito desta maneira:
```
N3 = ( APOL1 + APOL2 + APOL3 + APOL4 + APOL5)  
     ----------------------------------------  
                      5  
```
  
Em seguida iremos obter o cálculo de média da Atividade Prática e da Prova Discursiva:  
  
```
N4 = ( AP x 4 + PD x 6 )
    ---------------------
             10  
```
  
Para obter a média total desta máteria, iremos calcular a nota de sua prova objetiva, com a  
as médias das apols, e a média da prova discursiva + prova prática, multiplicando por seus respectivos pesos.  
  
```
MD = ( PO x 30 + N3 x 20 + N4 x 50 )
    ---------------------------------
                 100  
```
  
License
-------

LGPL-3.0

Author Information
------------------

Criado em 2018 por [isca](https://isca.space)

Enjoy :smile:
  
