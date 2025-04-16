# Aprendiendo GO

**Estructura de un programa Go**

Primero, definimos un paquete como `main`.

```go
package main
```

Entonces tenemos algunas importaciones

```go
import "fmt”
```

Por último, pero no menos importante, es nuestro `main` función que actúa como punto de entrada para nuestra aplicación, al igual que en otros lenguajes como C, Java o C#

```go
func main() {
	fmt.Println("Hello World!")
}
```

Ejecutar codigo con:

```go
$ go run main.go
Hello World!
```

## **Variables y Tipos de Datos**

Esto también se conoce como declaración sin inicialización:

```go
var foo string
```

Declaración de variables sin inicializar

```go
var foo string
```

Declaración de variables inicializadas

```go
var foo string = "go is awsome”
```

Múltiples declaraciones

```go
var foo, bar string = "Hello", "World”
o
var (
		foo string = "Hello"
		bar string  = "World"
	)
```

El tipo se omite pero se inferirá

Declaración abreviada, aquí omitimos `var` la palabra clave y el tipo siempre están implícitos. Así es como veremos las variables declaradas la mayor parte del tiempo. También usamos el `:=` para declaración más asignación

```go
foo := "Shorthand!”
```

### **Constantes**

También podemos declarar constantes con el `const` palabra clave. Que como su nombre indica, son valores fijos que no se pueden reasignar

```go
const constant = "This is a constant”
```

También es importante tener en cuenta que, solo las constantes se pueden asignar a otras constantes.

```go
const a = 10
const b = a // ✅ Works

var a = 10
const b = a // ❌ a (variable of type int) is not constant (InvalidConstInit)
```

### **Tipos de Datos**

### Cadenas

En Go, una cadena es una secuencia de bytes. Se declaran usando comillas dobles o backticks que pueden abarcar múltiples líneas

```go
var name string = "My name is Go"
var bio string = `I am statically typed.
											 									I was designed at Google.`
```

### Bool

```go
var value bool = false
var isItTrue bool = true
```

### **Operadores**

| **Tipo** | **Sintaxis** |
| -------- | ------------ | --- | ----- |
| Lógico   | `&&` `       |     | ` `!` |
| Igualdad | `==` `!=`    |

### **Tipos numéricos**

**Números enteros firmados y no firmados**

Go tiene varios tipos de enteros integrados de diferentes tamaños para almacenar enteros firmados y no firmados

El tamaño del genérico `int` y `uint` los tipos dependen de la plataforma. Esto significa que tiene 32 bits de ancho en un sistema de 32 bits y 64 bits de ancho en un sistema de 64 bits.

Son utiles con el uso eficiente de memoria, en aplicaciones donde la memoria es limitada, usar tipos más pequeños como **`int8`** o **`uint8`** puede ser crucial para optimizar el uso de recursos

```go
var i int = 404                     // Platform dependent
var i8 int8 = 127                   // -128 to 127
var i16 int16 = 32767               // -2^15 to 2^15 - 1
var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
```

Similar a los enteros firmados, tenemos enteros sin firmar.

```go
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64
var uiptr uintptr                     // Integer representation of a memory address
```

Si se dio cuenta, también hay un puntero entero sin firmar `uintptr` tipo, que es una representación entera de una dirección de memoria. No se recomienda usar esto, por lo que no tenemos que preocuparnos por ello.

**Entonces, ¿cuál deberíamos usar?**

Se recomienda que siempre que necesitemos un valor entero, solo debemos usar `int` a menos que tengamos una razón específica para usar un tipo entero de tamaño o sin firmar.

### **Byte y Rune**

Golang tiene dos tipos de enteros adicionales llamados `byte` y `rune` eso son alias para `uint8` y `int32` tipos de datos respectivamente.

Byte: Es un alias para el tipo **`uint8`**, lo que significa que representa un entero sin signo de 8 bits. Su rango va de 0 a 255

Se utiliza para representar datos binarios o caracteres en formato ASCII (que ocupan un solo byte)

```go
var b byte = 'A' // Representa el carácter 'A' en ASCII (valor 65)
fmt.Printf("Valor: %d, Carácter: %c\n", b, b)

```

Rune

un alias para int32, lo que significa que representa un entero con signo de 32 bits. Está diseñado para almacenar puntos de código Unicode, que pueden representar caracteres de cualquier idioma o símbolo

```go
var r rune = '❤' // Representa el carácter '❤' (Unicode U+2764)
fmt.Printf("Valor Unicode: %U, Carácter: %c\n", r, r)
```

```go
type byte = uint8
type rune = int32
```

_A `rune` representa un punto de código unicode._

```go
var b byte = 'a'
var r rune = '🍕'
```

**Punto flotante**

A continuación, tenemos tipos de punto flotante que se utilizan para almacenar números con un componente decimal.

Go tiene dos tipos de punto flotante `float32` y `float64`. Ambos tipos siguen el estándar IEEE-754.

_El tipo predeterminado para los valores de punto flotante es float64._

```go
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit
```

**Operadores**

Go proporciona varios operadores para realizar operaciones en tipos numéricos.

| **Tipo**              | **Sintaxis**                                    |
| --------------------- | ----------------------------------------------- | --------------- |
| Aritmética            | `+` `-` `*` `/` `%`                             |
| Comparación           | `==` `!=` `<` `>` `<=` `>=`                     |
| Bitwise               | `&` `                                           | ` `^` `<<` `>>` |
| Incremento/Decremento | `++` `--`                                       |
| Asignación            | `=` `+=` `-=` `*=` `/=` `%=` `<<=` `>>=` `&=` ` | =` `^=`         |

### Complex

Hay 2 tipos complex en Go, Go `complex128` donde están las partes reales e imaginarias `float64` y `complex64` donde las partes reales e imaginarias están `float32`.

```go
var c1 complex128 = complex(10, 1)
var c2 complex64 = 12 + 4i

```

Su uso puede ser menos común en aplicaciones generales de programación, ya que no todos los problemas requieren representar cantidades complejas

### **Cero Valores**

Ahora discutamos los valores cero. Entonces, en Go, cualquier variable declarada sin un valor inicial explícito recibe su *valor cero*. Por ejemplo, declaremos algunas variables y veamos

```go
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

Entonces, como podemos ver `int` y `float` se asignan como 0, `bool` como falso, y `string` como una cadena vacía. Esto es bastante diferente de cómo lo hacen otros idiomas. Por ejemplo, la mayoría de los idiomas inicializan variables no asignadas como nulas o indefinidas.

Esto es genial, pero ¿cuáles son esos símbolos porcentuales en nuestro `Printf` ¿función? Como ya has adivinado, se utilizan para formatear y aprenderemos sobre ellos más adelante.

### **Tipo Conversión**

Continuando, ahora que hemos visto cómo funcionan los tipos de datos, veamos cómo hacer la conversión de tipos.

```go
i := 42
f := float64(i)
u := uint(f)
fmt.Printf("%T %T", f, u)
```

Y como podemos ver, imprime el tipo como `float64` y `uint`.

### **Tipos de alias**

Los tipos de alias se introdujeron en Go 1.9. Permiten a los desarrolladores proporcionar un nombre alternativo para un tipo existente y usarlo indistintamente con el tipo subyacente.

Esto **no crea un nuevo tipo**, simplemente **le da otro nombre** a un tipo ya existente.

Es como decir:

> "Para este código, cada vez que diga MyAlias, en realidad me refiero a string."

```go
package main

import "fmt"

type MyAlias = string

func main() {
	var str MyAlias = "I am an alias"

	fmt.Printf("%T - %s", str, str) // Output: string - I am an alias
}
```

### **Tipos definidos**

Por último, hemos definido tipos que, a diferencia de los tipos de alias, no utilizan un signo igual.

```go
package main

import "fmt"

type MyDefined string

func main() {
	var str MyDefined = "I am defined"

	fmt.Printf("%T - %s", str, str) // Output: main.MyDefined - I am defined
}
```

Aquí **sí se crea un nuevo tipo** basado en otro (en este caso, basado en `string`), **pero es un tipo distinto**.

Es como decir:

> "Quiero un tipo nuevo que se comporte como un string, pero no es un string."

**Pero espera...¿cuál es la diferencia?**

Por lo tanto, los tipos definidos hacen más que solo dar un nombre a un tipo.

Primero define un nuevo tipo con nombre con un tipo subyacente. Sin embargo, este tipo definido es diferente de cualquier otro tipo, incluido su tipo subrayado.

Por lo tanto, no se puede usar indistintamente con los tipos de alias tipo subyacente.

Es un poco confuso al principio, con suerte, este ejemplo dejará las cosas claras.

```go
package main

import "fmt"

type MyAlias = string

type MyDefined string

func main() {
	var alias MyAlias
	var def MyDefined

	// ✅ Works
	var copy1 string = alias

	// ❌ Cannot use def (variable of type MyDefined) as string value in variable
	var copy2 string = def

	fmt.Println(copy1, copy2)
}
```

**🔍 ¿Por qué existe esta diferencia?**

- Los **alias** se usan cuando quieres renombrar un tipo por claridad o compatibilidad.
  - Ejemplo real: en paquetes donde quieres mantener compatibilidad con otra versión.
- Los **tipos definidos** se usan cuando quieres **crear tipos más seguros y específicos**, aunque internamente usen el mismo tipo base.
  - Ejemplo: podrías tener un tipo `UserID string` y `Email string`. Ambos son `string`, pero los defines por separado para evitar errores de asignación entre ellos.

## **Formato de Cuerda**

`fmt` el paquete contiene muchas funciones. Entonces, para ahorrar tiempo, discutiremos las funciones más utilizadas. Empecemos con `fmt.Print` dentro de nuestra función principal.

```go
fmt.Print("What", "is", "your", "name?")
fmt.Print("My", "name", "is", "golang")
```

Como podemos ver, `Print` no formatea nada, simplemente toma una cadena y la imprime.

A continuación, tenemos `Println` que es lo mismo que `Print` pero agrega una nueva línea al final y también inserta espacio entre los argumentos.

```go
fmt.Println("What", "is", "your", "name?")
fmt.Println("My", "name", "is", "golang")

```

A continuación, tenemos `Printf` también conocido como *"Formato de Impresión"*, lo que nos permite formatear números, cadenas, booleanos y mucho más.

```go
name := "golang"
fmt.Println("What is your name?")
fmt.Printf("My name is %s", name)
```

Como podemos ver eso `%s` fue sustituido con nuestro `name` variable.

Pero la pregunta es qué es `%s` ¿y qué significa?

Entonces, estos se llaman *verbos de anotación* y le dicen a la función cómo formatear los argumentos. Podemos controlar cosas como ancho, tipos y precisión con estos y hay muchos de ellos. Aquí hay un [hoja de trucos](https://pkg.go.dev/fmt).

Ahora, veamos rápidamente algunos ejemplos más. Aquí intentaremos calcular un porcentaje e imprimirlo en la consola.

### 🎯 Verbos de formato más comunes

| Verbo | Descripción                             | Ejemplo                               |
| ----- | --------------------------------------- | ------------------------------------- |
| `%d`  | Entero en base 10                       | `fmt.Printf("%d", 42)`                |
| `%f`  | Número de punto flotante decimal        | `fmt.Printf("%f", 3.14)`              |
| `%s`  | Cadena de texto                         | `fmt.Printf("%s", "hola")`            |
| `%v`  | Valor por defecto (genérico)            | `fmt.Printf("%v", anyValue)`          |
| `%T`  | Tipo del valor                          | `fmt.Printf("%T", "hola")` → `string` |
| `%q`  | Cadena entre comillas (útil para texto) | `fmt.Printf("%q", "hola")` → `"hola"` |
| `%x`  | Número hexadecimal (minúsculas)         | `fmt.Printf("%x", 255)` → `ff`        |
| `%X`  | Número hexadecimal (mayúsculas)         | `fmt.Printf("%X", 255)` → `FF`        |
| `%t`  | Booleano (`true` o `false`)             | `fmt.Printf("%t", true)`              |
| `%p`  | Dirección de memoria (puntero)          | `fmt.Printf("%p", &x)`                |

Digamos que solo queremos `77.78` que es la precisión de 2 puntos, podemos hacer eso también usando `.2f`.

Además, para agregar un signo de porcentaje real, tendremos que hacerlo *escápate de él*.

```go
percent := (7.0 / 9) * 100
fmt.Printf("%.2f %%", percent)
```

Esto nos lleva a `Sprint`, `Sprintln`, y `Sprintf`. Estas son básicamente las mismas que las funciones de impresión, la única diferencia es que devuelven la cadena en lugar de imprimirla

```go
s := fmt.Sprintf("hex:%x bin:%b", 10 ,10)
fmt.Println(s)
```

Entonces, como podemos ver `Sprintf` formatea nuestro entero como hexadecimal o binario y lo devuelve como una cadena.

Por último, tenemos literales de cuerda multilínea, que se pueden usar así.

```go
msg := `
Hello from
multiline
`

fmt.Println(msg)
```

¡Genial! Pero esto es solo la punta del iceberg...así que asegúrese de revisar el documento de go para `fmt` paquete.

Para aquellos que vienen de fondo C/C++, esto debería sentirse natural, pero si vienes de, digamos Python o JavaScript, esto podría ser un poco extraño al principio. Pero es muy potente y verás que esta funcionalidad se usa bastante ampliamente.

## **Control de Flujo**

### **If/Else**

Esto funciona más o menos de lo mismo que esperas, pero la expresión no necesita estar rodeada de paréntesis `()`.

```go
func main() {
x := 10
if x > 5 {
	fmt.Println("x is gt 5")
} else if x > 10 {
	fmt.Println("x is gt 10")
} else {
	fmt.Println("else case")
}
}
```

### **Compact If**

```go
func main() {
	if x := 10; x > 5 {
		fmt.Println("x is gt 5")
	}
}
```

### Switch

A continuación, tenemos `switch` declaración, que a menudo es una forma más corta de escribir lógica condicional.

En Go, el caso de cambio solo ejecuta el primer caso cuyo valor es igual a la expresión de condición y no todos los casos que siguen. Por lo tanto, a diferencia de otros idiomas, `break` la declaración se agrega automáticamente al final de cada caso.

Esto significa que evalúa los casos de arriba a abajo, deteniéndose cuando un caso tiene éxito. Echemos un vistazo a un ejemplo:

```go
func main() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}
}
```

Switch también admite una declaración abreviada como esta.

```go
switch day := "monday"; day {
case "monday":
	fmt.Println("time to work!")
case "friday":
	fmt.Println("let's party")
default:
	fmt.Println("browse memes")
}

```

También podemos usar el `fallthrough` palabra clave para transferir el control al siguiente caso a pesar de que el caso actual podría haber coincidido.

```go
switch day := "monday"; day {
case "monday":
	fmt.Println("time to work!")
	fallthrough
case "friday":
	fmt.Println("let's party")
default:
	fmt.Println("browse memes")
}

```

Y si ejecutamos esto, veremos que después de que el primer caso coincida, la instrucción del cambio continúa con el siguiente caso debido a la `fallthrough` palabra clave.

`go run main.go
time to work!
let's party`

También podemos usarlo sin ninguna condición, que es la misma que `switch true`.

```go
x := 10
switch {
case x > 5:
fmt.Println("x is greater")
default:
fmt.Println("x is not greater")
}
```

### Bucle

Ahora, volvamos nuestra atención hacia los bucles.

Así que en Go, solo tenemos un tipo de bucle que es el `for` bucle.

Pero es increíblemente versátil. Igual que si la declaración, para el bucle, no necesita ningún paréntesis `()` a diferencia de otros idiomas.

```go
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

Lo básico `for` loop tiene tres componentes separados por punto y coma:

- **declaración init**
  : que se ejecuta antes de la primera iteración.
- **expresión de condición**
  : que se evalúa antes de cada iteración.
- **declaración de correo**
  : que se ejecuta al final de cada iteración.

**Romper y continuar**

Como se esperaba, Go también es compatible con ambos `break` y `continue` declaraciones para el control de bucle. Probemos un ejemplo rápido:

```go
func main() {
	for i := 0; i < 10; i++ {
		if i < 2 {
			continue
		}

		fmt.Println(i)

		if i > 5 {
			break
		}
	}

	fmt.Println("We broke out!")
}
```

Entonces, el `continue` declaración se utiliza cuando queremos omitir la parte restante del bucle, y `break` la declaración se usa cuando queremos salir del bucle.

Además, las declaraciones de inicio y publicación son opcionales, por lo tanto, podemos hacer nuestras `for` el bucle también se comporta como un bucle de tiempo.

```go
func main() {
	i := 0

	for ;i < 10; {
		i += 1
	}
}
```

_Nota: también podemos eliminar los semicolones adicionales para que sea un poco más limpio._

Por último, si omitimos la condición de bucle, se bucle para siempre, por lo que un bucle infinito se puede expresar de forma compacta. Esto también se conoce como el bucle para siempre.

```go
func main() {
for {
// do stuff here
}
}
```

## Funciones

**Declaración simple**

```go
func myFunction() {}
```

Y podemos *llamar o ejecutar* es como sigue.

```go
myFunction()
```

Pasemos algunos parámetros a ello.

```go
func main() {
	myFunction("Hello")
}

func myFunction(p1 string) {
	fmt.Println(p1)
}
```

Ahora también devolvamos un valor.

```go
func main() {
	s := myFunction("Hello")
	fmt.Println(s)
}

func myFunction(p1 string) string {
	msg := fmt.Sprintf("%s function", p1)
	return msg
}
```

### **Devoluciones nombradas**

Otra característica genial es [devoluciones nombradas](https://go.dev/tour/basics/7), donde los valores de retorno pueden ser nombrados y tratados como sus propias variables.

Las **devoluciones nombradas** (_named returns_) en Go son una característica que te permite **nombrar las variables de retorno directamente en la firma de la función**, como parte del encabezado. Esto puede ayudarte a escribir código más claro y a veces evitar declarar variables adicionales dentro de la función.

```go
func dividir(a, b float64) (resultado float64, err error) {
    if b == 0 {
        err = fmt.Errorf("no se puede dividir entre cero")
        return // usa las variables nombradas: resultado, err
    }
    resultado = a / b
    return // return implícito
}
```

**Equivalente sin nombres:**

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir entre cero")
    }
    return a / b, nil
}
```

Diré que, aunque esta característica es interesante, úsela con cuidado, ya que esto podría reducir la legibilidad para funciones más grandes.

### **Funciones como valores**

A continuación, hablemos de funciones como valores, en Go las funciones son de primera clase y podemos utilizarlas como valores. Entonces, ¡limpiemos nuestra función y probémosla!

En Go, **las funciones son valores de primera clase**. Eso significa que puedes tratarlas como cualquier otro valor: **asignarlas a variables, pasarlas como argumentos a otras funciones o devolverlas como resultado**. A esto se le llama trabajar con **funciones como valores** (_functions as values_).

- **Asignar una función a una variable**
- **Pasar una función como parámetro**
- **Devolver una función desde otra función**

```go
func myFunction() {
	fn := func() {
		fmt.Println("inside fn")
	}

	fn()
}
```

También podemos simplificar esto haciendo `fn` un *función anónima*.

```go
func myFunction() {
	func() {
		fmt.Println("inside fn")
	}()
}
```

```go
package main

import "fmt"

func saludar(nombre string) {
    fmt.Println("Hola", nombre)
}

func main() {
    var f func(string) // declarar variable de tipo función
    f = saludar        // asignar función a la variable
    f("Gopher")        // llamar la función a través de la variable
}

```

### **Closures**

¿Por qué parar allí? también devolvamos una función y, por lo tanto, creemos algo llamado cierre. Una definición simple puede ser que un cierre es un valor de función que hace referencia a variables de fuera de su cuerpo.

Los cierres se exploran léxicamente, lo que significa que las funciones pueden acceder a los valores en alcance al definir la función.

> Un closure "recuerda" las variables del entorno donde fue definido, aunque ese entorno ya no exista.

**🧪 ¿Por qué es útil?**

1. ✅ Puedes mantener **estado interno** sin variables globales.
2. ✅ Ideal para **generadores, contadores, filtros dinámicos**, etc.
3. ✅ Puedes crear **funciones con configuración personalizada**.

```go
func myFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}

add := myFunction()

add(5)
fmt.Println(add(10))
```

Como podemos ver, obtenemos un resultado de 15 como `sum` la variable es *atado* a la función. Este es un concepto muy poderoso y definitivamente, un conocimiento imprescindible.

### **Funciones Variádicas**

Ahora veamos las funciones variádicas, que son funciones que pueden tomar cero o múltiples argumentos usando el `...` operador de elipses.

Un ejemplo aquí sería una función que puede agregar un montón de valores.

```go
func main() {
	sum := add(1, 2, 3, 5)
	fmt.Println(sum)
}

func add(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}
```

### Init

En Go, `init` es una función especial del ciclo de vida que se ejecuta antes de la `main` función.

Similar a `main`, el `init` la función no toma ningún argumento ni devuelve ningún valor. Veamos cómo funciona con un ejemplo.

```go
package main

import "fmt"

func init() {
	fmt.Println("Before main!")
}

func main() {
	fmt.Println("Running main")
}
```

Como se esperaba, el `init` la función se ejecutó antes de la `main` función.

A diferencia `main`, puede haber más de uno `init` función en archivos individuales o múltiples.

Para múltiples `init` en un solo archivo, su procesamiento se realiza en el orden de su declaración, mientras `init` las funciones declaradas en múltiples archivos se procesan de acuerdo con el orden de nombre de archivo lexicográfico.

```go
package main

import "fmt"

func init() {
	fmt.Println("Before main!")
}

func init() {
	fmt.Println("Hello again?")
}

func main() {
	fmt.Println("Running main")
}
```

Y si corremos esto, veremos el `init` las funciones se ejecutaron en el orden en que fueron declaradas.

El `init` la función es opcional y se utiliza particularmente para cualquier configuración global que pueda ser esencial para nuestro programa, como establecer una conexión de base de datos, buscar archivos de configuración, configurar variables de entorno, etc.

### **Defer**

Por último, discutamos el `defer` palabra clave, que nos permite posponer la ejecución de una función hasta que la función circundante regrese.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("I am finished")
	fmt.Println("Doing some work...")
}
```

¿Podemos usar múltiples funciones de aplazamiento? Absolutamente, esto nos lleva a lo que se conoce como *aplazar pila*. Echemos un vistazo a un ejemplo:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")

	fmt.Println("Doing some work...")
}

go run main.go
Doing some work...
Are you?
I am finished
```

Como podemos ver, las declaraciones de aplazamiento se apilan y ejecutan en un *último en salir* manera.

Por lo tanto, Defer es increíblemente útil y se usa comúnmente para hacer limpieza o manejo de errores.

Las funciones también se pueden usar con genéricos, pero las discutiremos más adelante en el curso.

## Modulos
