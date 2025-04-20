# Aprendiendo GO

# Capitulo 2

## **Estructura de un programa Go**

Primero, definimos un paquete como `main`.

```go
package main
```

Entonces tenemos algunas importaciones:

```go
import "fmt"
```

Por último, pero no menos importante, está nuestra función `main` que actúa como punto de entrada para nuestra aplicación, al igual que en otros lenguajes como C, Java o C#:

```go
func main() {
    fmt.Println("Hello World!")
}
```

Ejecutar código con:

```bash
$ go run main.go
Hello World!
```

## **Variables y Tipos de Datos**

### **Declaración de variables**

Declaración sin inicialización:

```go
var foo string
```

Declaración de variables inicializadas:

```go
var foo string = "go is awesome"
```

Múltiples declaraciones:

```go
var foo, bar string = "Hello", "World"
// o
var (
    foo string = "Hello"
    bar string = "World"
)
```

Declaración abreviada, aquí omitimos la palabra clave `var` y el tipo siempre está implícito. Así es como veremos las variables declaradas la mayor parte del tiempo. También usamos el `:=` para declaración más asignación:

```go
foo := "Shorthand!"
```

### **Constantes**

También podemos declarar constantes con la palabra clave `const`. Como su nombre indica, son valores fijos que no se pueden reasignar:

```go
const constant = "This is a constant"
```

También es importante tener en cuenta que solo las constantes se pueden asignar a otras constantes:

```go
const a = 10
const b = a // ✅ Funciona

var a = 10
const b = a // ❌ a (variable de tipo int) no es constante (InvalidConstInit)
```

### **Tipos de Datos**

#### **Cadenas**

En Go, una cadena es una secuencia de bytes. Se declaran usando comillas dobles o backticks que pueden abarcar múltiples líneas:

```go
var name string = "My name is Go"
var bio string = `I am statically typed.
I was designed at Google.`
```

#### **Bool**

```go
var value bool = false
var isItTrue bool = true
```

#### **Operadores**

| **Tipo** | **Sintaxis**    |
| -------- | --------------- |
| Lógico   | `&&` `\|\|` `!` |
| Igualdad | `==` `!=`       |

### **Tipos numéricos**

#### **Números enteros firmados y no firmados**

Go tiene varios tipos de enteros integrados de diferentes tamaños para almacenar enteros firmados y no firmados.

El tamaño del genérico `int` y `uint` depende de la plataforma. Esto significa que tiene 32 bits de ancho en un sistema de 32 bits y 64 bits de ancho en un sistema de 64 bits.

Son útiles para el uso eficiente de memoria. En aplicaciones donde la memoria es limitada, usar tipos más pequeños como **`int8`** o **`uint8`** puede ser crucial para optimizar el uso de recursos.

```go
var i int = 404                     // Dependiente de la plataforma
var i8 int8 = 127                   // -128 a 127
var i16 int16 = 32767               // -2^15 a 2^15 - 1
var i32 int32 = -2147483647         // -2^31 a 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 a 2^63 - 1
```

Similar a los enteros firmados, tenemos enteros sin firmar:

```go
var ui uint = 404                     // Dependiente de la plataforma
var ui8 uint8 = 255                   // 0 a 255
var ui16 uint16 = 65535               // 0 a 2^16
var ui32 uint32 = 2147483647          // 0 a 2^32
var ui64 uint64 = 9223372036854775807 // 0 a 2^64
var uiptr uintptr                     // Representación entera de una dirección de memoria
```

Si te has dado cuenta, también hay un puntero entero sin firmar `uintptr`, que es una representación entera de una dirección de memoria. No se recomienda usar esto, por lo que no tenemos que preocuparnos por ello.

**Entonces, ¿cuál deberíamos usar?**

Se recomienda que siempre que necesitemos un valor entero, solo debemos usar `int` a menos que tengamos una razón específica para usar un tipo entero de tamaño específico o sin firmar.

### **Byte y Rune**

Golang tiene dos tipos de enteros adicionales llamados `byte` y `rune` que son alias para los tipos de datos `uint8` y `int32` respectivamente:

```go
type byte = uint8
type rune = int32
```

**Byte**: Es un alias para el tipo **`uint8`**, lo que significa que representa un entero sin signo de 8 bits. Su rango va de 0 a 255.

Se utiliza para representar datos binarios o caracteres en formato ASCII (que ocupan un solo byte):

```go
var b byte = 'A' // Representa el carácter 'A' en ASCII (valor 65)
fmt.Printf("Valor: %d, Carácter: %c\n", b, b)
```

**Rune**: Es un alias para int32, lo que significa que representa un entero con signo de 32 bits. Está diseñado para almacenar puntos de código Unicode, que pueden representar caracteres de cualquier idioma o símbolo:

```go
var r rune = '❤' // Representa el carácter '❤' (Unicode U+2764)
fmt.Printf("Valor Unicode: %U, Carácter: %c\n", r, r)
```

Un `rune` representa un punto de código unicode:

```go
var b byte = 'a'
var r rune = '🍕'
```

### **Punto flotante**

A continuación, tenemos tipos de punto flotante que se utilizan para almacenar números con un componente decimal.

Go tiene dos tipos de punto flotante: `float32` y `float64`. Ambos tipos siguen el estándar IEEE-754.

_El tipo predeterminado para los valores de punto flotante es float64._

```go
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit
```

### **Operadores numéricos**

Go proporciona varios operadores para realizar operaciones en tipos numéricos:

| **Tipo**              | **Sintaxis**                                             |
| --------------------- | -------------------------------------------------------- |
| Aritmética            | `+` `-` `*` `/` `%`                                      |
| Comparación           | `==` `!=` `<` `>` `<=` `>=`                              |
| Bitwise               | `&` `\|` `^` `<<` `>>`                                   |
| Incremento/Decremento | `++` `--`                                                |
| Asignación            | `=` `+=` `-=` `*=` `/=` `%=` `<<=` `>>=` `&=` `\|=` `^=` |

### **Complex**

Hay 2 tipos complex en Go: `complex128` donde las partes reales e imaginarias son `float64` y `complex64` donde las partes reales e imaginarias son `float32`:

```go
var c1 complex128 = complex(10, 1)
var c2 complex64 = 12 + 4i
```

Su uso puede ser menos común en aplicaciones generales de programación, ya que no todos los problemas requieren representar cantidades complejas.

### **Cero Valores**

En Go, cualquier variable declarada sin un valor inicial explícito recibe su _valor cero_. Por ejemplo:

```go
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

Como podemos ver, `int` y `float` se asignan como 0, `bool` como falso, y `string` como una cadena vacía. Esto es diferente de cómo lo hacen otros idiomas. Por ejemplo, la mayoría de los idiomas inicializan variables no asignadas como nulas o indefinidas.

Esto es genial, pero ¿qué son esos símbolos porcentuales en nuestra función `Printf`? Como ya has adivinado, se utilizan para formatear y aprenderemos sobre ellos más adelante.

### **Conversión de Tipos**

Ahora que hemos visto cómo funcionan los tipos de datos, veamos cómo hacer la conversión de tipos:

```go
i := 42
f := float64(i)
u := uint(f)
fmt.Printf("%T %T", f, u)
```

Y como podemos ver, imprime los tipos como `float64` y `uint`.

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

Por último, tenemos los tipos definidos que, a diferencia de los tipos de alias, no utilizan un signo igual:

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

Los tipos definidos hacen más que solo dar un nombre a un tipo.

Primero define un nuevo tipo con nombre con un tipo subyacente. Sin embargo, este tipo definido es diferente de cualquier otro tipo, incluido su tipo subrayado.

Por lo tanto, no se puede usar indistintamente con los tipos de alias o su tipo subyacente.

Es un poco confuso al principio, con suerte, este ejemplo dejará las cosas claras:

```go
package main

import "fmt"

type MyAlias = string

type MyDefined string

func main() {
    var alias MyAlias
    var def MyDefined

    // ✅ Funciona
    var copy1 string = alias

    // ❌ No se puede usar def (variable de tipo MyDefined) como valor string en variable
    var copy2 string = def

    fmt.Println(copy1, copy2)
}
```

**🔍 ¿Por qué existe esta diferencia?**

- Los **alias** se usan cuando quieres renombrar un tipo por claridad o compatibilidad.
  - Ejemplo real: en paquetes donde quieres mantener compatibilidad con otra versión.
- Los **tipos definidos** se usan cuando quieres **crear tipos más seguros y específicos**, aunque internamente usen el mismo tipo base.
  - Ejemplo: podrías tener un tipo `UserID string` y `Email string`. Ambos son `string`, pero los defines por separado para evitar errores de asignación entre ellos.

## **Formato de Cadenas**

El paquete `fmt` contiene muchas funciones. Discutiremos las funciones más utilizadas. Empecemos con `fmt.Print` dentro de nuestra función principal:

```go
fmt.Print("What", "is", "your", "name?")
fmt.Print("My", "name", "is", "golang")
```

Como podemos ver, `Print` no formatea nada, simplemente toma una cadena y la imprime.

A continuación, tenemos `Println` que es lo mismo que `Print` pero agrega una nueva línea al final y también inserta espacio entre los argumentos:

```go
fmt.Println("What", "is", "your", "name?")
fmt.Println("My", "name", "is", "golang")
```

A continuación, tenemos `Printf` también conocido como _"Formato de Impresión"_, lo que nos permite formatear números, cadenas, booleanos y mucho más:

```go
name := "golang"
fmt.Println("What is your name?")
fmt.Printf("My name is %s", name)
```

Como podemos ver, `%s` fue sustituido con nuestra variable `name`.

Pero la pregunta es qué es `%s` y qué significa?

Son _verbos de anotación_ y le dicen a la función cómo formatear los argumentos. Podemos controlar cosas como ancho, tipos y precisión con estos y hay muchos de ellos. Aquí hay una [hoja de trucos](https://pkg.go.dev/fmt).

Veamos rápidamente algunos ejemplos más. Aquí intentaremos calcular un porcentaje e imprimirlo en la consola.

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

Digamos que solo queremos `77.78`, que es la precisión de 2 decimales, podemos hacer eso también usando `.2f`.

Además, para agregar un signo de porcentaje real, tendremos que escaparlo:

```go
percent := (7.0 / 9) * 100
fmt.Printf("%.2f %%", percent)
```

Esto nos lleva a `Sprint`, `Sprintln`, y `Sprintf`. Estas son básicamente las mismas que las funciones de impresión, la única diferencia es que devuelven la cadena en lugar de imprimirla:

```go
s := fmt.Sprintf("hex:%x bin:%b", 10, 10)
fmt.Println(s)
```

Como podemos ver, `Sprintf` formatea nuestro entero como hexadecimal o binario y lo devuelve como una cadena.

Por último, tenemos literales de cadena multilínea, que se pueden usar así:

```go
msg := `
Hello from
multiline
`

fmt.Println(msg)
```

¡Genial! Pero esto es solo la punta del iceberg...así que asegúrese de revisar la documentación de Go para el paquete `fmt`.

Para aquellos que vienen del fondo C/C++, esto debería sentirse natural, pero si vienes de Python o JavaScript, esto podría ser un poco extraño al principio. Pero es muy potente y verás que esta funcionalidad se usa bastante ampliamente.

## **Control de Flujo**

### **If/Else**

Esto funciona más o menos de lo mismo que esperas, pero la expresión no necesita estar rodeada de paréntesis `()`:

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

### **Switch**

A continuación, tenemos la declaración `switch`, que a menudo es una forma más corta de escribir lógica condicional.

En Go, el caso de switch solo ejecuta el primer caso cuyo valor es igual a la expresión de condición y no todos los casos que siguen. Por lo tanto, a diferencia de otros idiomas, la declaración `break` se agrega automáticamente al final de cada caso.

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

Switch también admite una declaración abreviada como esta:

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

También podemos usar la palabra clave `fallthrough` para transferir el control al siguiente caso a pesar de que el caso actual podría haber coincidido:

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

Y si ejecutamos esto, veremos que después de que el primer caso coincida, la instrucción del switch continúa con el siguiente caso debido a la palabra clave `fallthrough`:

```
time to work!
let's party
```

También podemos usarlo sin ninguna condición, que es lo mismo que `switch true`:

```go
x := 10
switch {
case x > 5:
    fmt.Println("x is greater")
default:
    fmt.Println("x is not greater")
}
```

### **Bucles**

En Go, solo tenemos un tipo de bucle que es el bucle `for`.

Pero es increíblemente versátil. Igual que en la declaración if, para el bucle, no necesita ningún paréntesis `()` a diferencia de otros idiomas:

```go
func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

El bucle `for` básico tiene tres componentes separados por punto y coma:

- **declaración init**: que se ejecuta antes de la primera iteración.
- **expresión de condición**: que se evalúa antes de cada iteración.
- **declaración post**: que se ejecuta al final de cada iteración.

**Break y Continue**

Como era de esperar, Go también es compatible con las declaraciones `break` y `continue` para el control de bucle. Probemos un ejemplo rápido:

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

La declaración `continue` se utiliza cuando queremos omitir la parte restante del bucle, y la declaración `break` se usa cuando queremos salir del bucle.

Además, las declaraciones de inicio y post son opcionales, por lo tanto, podemos hacer que nuestro bucle `for` también se comporte como un bucle while:

```go
func main() {
    i := 0

    for i < 10 {
        i += 1
    }
}
```

_Nota: también podemos eliminar los semicolones adicionales para que sea un poco más limpio._

Por último, si omitimos la condición de bucle, se bucle para siempre, por lo que un bucle infinito se puede expresar de forma compacta. Esto también se conoce como el bucle para siempre:

```go
func main() {
    for {
        // hacer cosas aquí
    }
}
```

## **Funciones**

### **Declaración simple**

```go
func myFunction() {}
```

Y podemos _llamarla o ejecutarla_ como sigue:

```go
myFunction()
```

Pasemos algunos parámetros:

```go
func main() {
    myFunction("Hello")
}

func myFunction(p1 string) {
    fmt.Println(p1)
}
```

Ahora también devolvamos un valor:

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

Otra característica genial es [devoluciones nombradas](https://go.dev/tour/basics/7), donde los valores de retorno pueden ser nombrados y tratados como sus propias variables.

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

Aunque esta característica es interesante, úsela con cuidado, ya que podría reducir la legibilidad para funciones más grandes.

### **Funciones como valores**

En Go, las funciones son de primera clase y podemos utilizarlas como valores:

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

También podemos simplificar esto haciendo de `fn` una _función anónima_:

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

También podemos devolver una función y, por lo tanto, crear algo llamado closure. Una definición simple puede ser que un closure es un valor de función que hace referencia a variables de fuera de su cuerpo.

Los closures se exploran léxicamente, lo que significa que las funciones pueden acceder a los valores en alcance al definir la función.

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

func main() {
    add := myFunction()

    add(5)
    fmt.Println(add(10)) // Imprime 15
}
```

Como podemos ver, obtenemos un resultado de 15 ya que la variable `sum` está _atada_ a la función. Este es un concepto muy poderoso y definitivamente un conocimiento imprescindible.

### **Funciones Variádicas**

Ahora veamos las funciones variádicas, que son funciones que pueden tomar cero o múltiples argumentos usando el operador de elipses `...`.

Un ejemplo sería una función que puede sumar un montón de valores:

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

### **Init**

En Go, `init` es una función especial del ciclo de vida que se ejecuta antes de la función `main`.

Similar a `main`, la función `init` no toma ningún argumento ni devuelve ningún valor. Veamos cómo funciona con un ejemplo:

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

Como era de esperar, la función `init` se ejecutó antes de la función `main`.

A diferencia de `main`, puede haber más de una función `init` en archivos individuales o múltiples.

Para múltiples `init` en un solo archivo, su procesamiento se realiza en el orden de su declaración, mientras que las funciones `init` declaradas en múltiples archivos se procesan de acuerdo con el orden del nombre de archivo lexicográfico:

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

Y si ejecutamos esto, veremos que las funciones `init` se ejecutaron en el orden en que fueron declaradas.

La función `init` es opcional y se utiliza particularmente para cualquier configuración global que pueda ser esencial para nuestro programa, como establecer una conexión de base de datos, buscar archivos de configuración, configurar variables de entorno, etc.

### **Defer**

Por último, discutamos la palabra clave `defer`, que nos permite posponer la ejecución de una función hasta que la función circundante regrese:

```go
package main

import "fmt"

func main() {
    defer fmt.Println("I am finished")
    fmt.Println("Doing some work...")
}
```

¿Podemos usar múltiples funciones de aplazamiento? Absolutamente, esto nos lleva a lo que se conoce como _pila de defer_. Echemos un vistazo a un ejemplo:

```go
package main

import "fmt"

func main() {
    defer fmt.Println("I am finished")
    defer fmt.Println("Are you?")

    fmt.Println("Doing some work...")
}
```

Resultado:

```
Doing some work...
Are you?
I am finished
```

Como podemos ver, las declaraciones de defer se apilan y ejecutan en una manera de _último en entrar, primero en salir_ (LIFO).

Por lo tanto, Defer es increíblemente útil y se usa comúnmente para hacer limpieza o manejo de errores.

Las funciones también se pueden usar con genéricos, pero las discutiremos más adelante.

# Módulos en Go

Simplemente definido, un módulo es una colección de [paquetes GO](https://go.dev/ref/spec#Packages) almacenados en un árbol de archivos con un archivo `go.mod` en su raíz, siempre que el directorio esté _fuera_ de `$GOPATH/src`.

Los módulos Go se introdujeron en Go 1.11, que brinda soporte nativo para versiones y módulos. Antes, necesitábamos la marca `GO111MODULE=on` para activar la funcionalidad de los módulos cuando era experimental. Pero ahora, después de Go 1.13, el modo de módulos es el predeterminado para todo el desarrollo.

## ¿Qué es `GOPATH`?

`GOPATH` es una variable que define la raíz de su espacio de trabajo y contiene las siguientes carpetas:

- **src**: contiene el código fuente de Go organizado en una jerarquía.
- **pkg**: contiene código de paquete compilado.
- **bin**: contiene binarios compilados y ejecutables.

## Creación de módulos

Creemos un nuevo módulo usando el comando `go mod init` que crea un nuevo módulo e inicializa el archivo `go.mod` que lo describe:

```go
go mod init example
```

Lo importante a tener en cuenta aquí es que un módulo Go también puede corresponder a un repositorio de GitHub si planea publicar este módulo.

## El archivo go.mod

Exploremos `go.mod`, que es el archivo que define la _ruta del módulo_ y también la ruta de importación utilizada para el directorio raíz, y sus _requisitos de dependencia_:

```go
module <name>
go <version>

require (
...
)
```

## Gestión de dependencias

Para añadir una nueva dependencia, usaremos el comando `go install` o `go get`:

```go
go install github.com/rs/zerolog
# o
go get github.com/rs/zerolog
```

Al añadir dependencias, también se crea un archivo `go.sum`. Este archivo contiene los [hashes](https://go.dev/cmd/go/#hdr-Module_downloading_and_verification) esperados del contenido de los nuevos módulos.

Podemos enumerar todas las dependencias utilizando el comando `go list`:

```go
go list -m all
```

Si una dependencia no se utiliza, simplemente podemos eliminarla usando el comando `go mod tidy`:

```go
go mod tidy
```

## Vendor

La vendorización es el acto de hacer su propia copia de los paquetes de terceros que está utilizando su proyecto. Esas copias se colocan tradicionalmente dentro de cada proyecto y luego se guardan en el repositorio del proyecto.

👉 **Copia todas las dependencias** (que están listadas en tu `go.mod`) dentro de una carpeta llamada `/vendor/`.

Esto se puede hacer a través del comando `go mod vendor`.

### ¿Para qué sirve?

1. ✅ **Entornos sin acceso a Internet**: puedes compilar sin necesidad de descargar paquetes de nuevo.
2. ✅ **Builds reproducibles**: garantiza que siempre usas la misma versión del código externo.
3. ✅ **Revisión de código**: algunas empresas prefieren revisar el código de las dependencias incluyéndolo en el repo.
4. ✅ **Despliegue en producción**: puedes subir todo tu proyecto (incluidas dependencias) a servidores sin conexión a Internet.

### Ejemplo de vendorización

```go
package main
import "github.com/rs/zerolog/log"

func main() {
  log.Info().Msg("Hello")
}
```

```
go mod tidy
go: finding module for package github.com/rs/zerolog/log
go: found github.com/rs/zerolog/log in github.com/rs/zerolog v1.26.1

go mod vendor
```

Después de ejecutar el comando `go mod vendor`, se creará el directorio `vendor`:

```
├── go.mod
├── go.sum
├── go.work
├── main.go
└── vendor
    ├── github.com
    │   └── rs
    │       └── zerolog
    │           └── ...
    └── modules.txt
```

# Paquetes

## ¿Qué son los paquetes?

Un paquete no es más que un directorio que contiene uno o más archivos de origen Go u otros paquetes Go.

Cada archivo de origen de Go debe pertenecer a un paquete y la declaración del paquete se realiza en la parte superior de cada archivo de origen de la siguiente manera:

```go
package <package_name>
```

El paquete `main` también debe contener una función `main()` que es una función especial que actúa como el punto de entrada de un programa ejecutable.

## Creando nuestro propio paquete

Echemos un vistazo a un ejemplo creando nuestro propio paquete `custom` y agregarle algunos archivos fuente, como `code.go`:

```go
package custom

var value int = 10 // No será exportado
var Value int = 20 // Será exportado
```

Los identificadores con minúsculas no se exportarán y serán privados para el paquete en el que se definen. Los identificadores con mayúsculas se exportarán y serán accesibles desde otros paquetes.

## Importando paquetes

Para importar y acceder a nuestro paquete `custom` en `main.go`:

```go
// go.mod
module example
go 1.18

// main.go
package main

import "example/custom"

func main() {
    // Acceso a Value (exportado)
    custom.Value
}
```

_Observe cómo el nombre del paquete es el último elemento de la ruta de importación._

También podemos importar múltiples paquetes:

```go
package main

import (
    "fmt"
    "example/custom"
)

func main() {
    fmt.Println(custom.Value)
}
```

Y podemos usar alias para nuestras importaciones para evitar colisiones:

```go
package main

import (
    "fmt"
    abcd "example/custom"
)

func main() {
    fmt.Println(abcd.Value)
}
```

## Dependencias Externas

En Go, no solo estamos limitados a trabajar con paquetes locales, también podemos instalar paquetes externos usando el comando `go install`:

```
$ go install github.com/rs/zerolog
```

```go
package main

import (
    "github.com/rs/zerolog/log"
    abcd "example/custom"
)

func main() {
    log.Print(abcd.Value)
}
```

Además, asegúrese de revisar la documentación de los paquetes que instala, que generalmente se encuentra en el archivo README del proyecto.

Go no tiene una convención particular de _"estructura de carpeta"_; siempre trate de organizar sus paquetes de una manera simple e intuitiva.

# Espacios de trabajo

Los espacios de trabajo de múltiples módulos se introdujeron en Go 1.18.

Los espacios de trabajo nos permiten trabajar con varios módulos simultáneamente sin tener que editar archivos `go.mod` para cada módulo. Cada módulo dentro de un espacio de trabajo se trata como un módulo raíz al resolver dependencias.

## Creando un espacio de trabajo

Comencemos creando un módulo `hello`:

```
mkdir workspaces && cd workspaces
mkdir hello && cd hello
go mod init hello
```

Para fines de demostración, añadiremos un simple `main.go`:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(Reverse("Hello, 世界!"))
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

Ahora, creemos nuestro espacio de trabajo en el directorio `workspaces`:

```
go work init
```

Esto creará un archivo `go.work`.

También agregaremos nuestro módulo `hello` al espacio de trabajo:

```
go work use ./hello
```

Esto actualizará el archivo `go.work` con referencia a nuestro módulo `hello`:

```
go 1.24.2

use ./hello
```

## Ejemplo con varios módulos

Podemos organizar mejor nuestro código moviendo la función `Reverse` a un módulo de utilidades:

**utils/reverse.go**:

```go
package utils

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

**hello/main.go**:

```go
package main

import (
    "fmt"
    "hello/utils"
)

func main() {
    fmt.Println(utils.Reverse("Hello, world!"))
}
```

Si ambos módulos tienen su propio archivo `go.mod`, **esto solo funciona** si tienes un archivo `go.work` que los una.

# Comandos Útiles

Durante nuestra discusión del módulo, mencionamos algunos comandos de Go relacionados con los módulos. Veamos otros comandos importantes:

## go fmt

Formatea el código fuente. Go enfatiza que debemos centrarnos en cómo debe funcionar nuestro código en lugar de cómo debe verse:

```
$ go fmt
```

Esto puede parecer extraño al principio, especialmente si vienes de JavaScript o Python, pero es bastante agradable no preocuparse por las reglas de formato.

## go vet

Informa de posibles errores en nuestros paquetes:

```
$ go vet
```

## go env

Imprime toda la información del entorno de Go:

```
$ go env
```

## go doc

Muestra la documentación de un paquete o símbolo:

```
$ go doc -src fmt Printf
```

## Otros comandos

Use el comando `go help` para ver qué otros comandos están disponibles:

```
$ go help
```

Algunos comandos adicionales importantes:

- `go fix`: encuentra programas Go que usan API antiguas y las reescribe para usar las más nuevas.
- `go generate`: se utiliza generalmente para la generación de código.
- `go install`: compila e instala paquetes y dependencias.
- `go clean`: se utiliza para limpiar archivos generados por compiladores.
- `go build`: compila paquetes y dependencias.
- `go test`: ejecuta pruebas.

## Build

Construir binarios estáticos es una de las mejores características de Go que nos permite enviar nuestro código de manera eficiente.

Podemos hacer esto muy fácilmente usando el `go build` comando.

```go
package main

import "fmt"

func main() {
	fmt.Println("I am a binary!")
}
```

Esto debería producir un binario con el nombre de nuestro módulo. Por ejemplo, aquí tenemos `example`.

También podemos especificar la salida.

```
$ go build -o app
```

Ahora, para ejecutar esto, simplemente necesitamos ejecutarlo.

`$ ./app
I am a binary!`

Ahora, hablemos de algunas variables importantes de tiempo de construcción, comenzando con:

- `GOOS`
- `GOARCH`

Estas variables de entorno nos ayudan a construir programas go para diferentes [sistemas operativos](https://en.wikipedia.org/wiki/Operating_system) y procesador subyacente [arquitecturas](https://en.wikipedia.org/wiki/Microarchitecture).

Podemos enumerar toda la arquitectura compatible utilizando `go tool` comando.

```
$ go tool dist list
android/amd64
ios/amd64
js/wasm
linux/amd64
windows/arm64
.
.
.
```

¡Aquí hay un ejemplo para construir un ejecutable de Windows desde macOS!

`$ GOOS=windows GOARCH=amd64 go build -o app.exe`

- `CGO_ENABLED`

Esta variable nos permite configurar [CGO](https://go.dev/blog/cgo), que es una forma de ir a llamar al código C.

Esto nos ayuda a producir un [binario estáticamente vinculado](https://en.wikipedia.org/wiki/Static_build) eso funciona sin dependencias externas.

Esto es bastante útil para, digamos cuando queremos ejecutar nuestros binarios go en un contenedor docker con dependencias externas mínimas.

Aquí hay un ejemplo de cómo usarlo:

`$ CGO_ENABLED=0 go build -o app`

# Capitulo 2

## Punteros

Simplemente definido, un puntero es una variable que se utiliza para almacenar la dirección de memoria de otra variable.

Se puede usar así:

```go
var x *T
```

Donde `T` es el tipo tal como `int`, `string`, `float`y así sucesivamente.

Probemos un ejemplo simple y viéndolo en acción.

```go
package main

import "fmt"

func main() {
	var p *int

	fmt.Println(p)
}
```

`$ go run main.go
nil`

Hmm, esto imprime `nil`, pero lo que es `nil`¿?

Así que nil es un identificador predeclarado en Go que representa el valor cero para punteros, interfaces, canales, mapas y cortes.

Esto es como lo que aprendimos en la sección de variables y tipos de datos, donde vimos eso sin inicializar `int` tiene un valor cero de 0, a `bool` tiene falso, y así sucesivamente.

Bien, ahora asignemos un valor al puntero.

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("address:", p)
}
```

Usamos el `&` operador de Ampersand para referirse a la dirección de memoria de una variable.

```
$ go run main.go
0xc0000b8000
```

Este debe ser el valor de la dirección de memoria de la variable `a`.

**Desreferenciación**

ambién podemos usar el `*` operador de asterisco para recuperar el valor almacenado en la variable a la que apunta el puntero. Esto también se llama **desreferenciación**.

Por ejemplo, podemos acceder al valor de la variable `a` a través del puntero `p` usando eso `*` operador de asterisco.

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("address:", p)
	fmt.Println("value:", *p)
}
```

`$ go run main.go
address: 0xc000018030
value: 10`

No solo podemos acceder a él, sino también cambiarlo a través del puntero.

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("before", a)
	fmt.Println("address:", p)

	*p = 20
	fmt.Println("after:", a)
}
```

```
$ go run main.go
before 10
address: 0xc000192000
after: 20
```

¡Creo que esto es bastante ordenado!

**Punteros como función args**

Los punteros también se pueden usar como argumentos para una función cuando necesitamos pasar algunos datos por referencia.

Aquí hay un ejemplo:

```go
myFunction(&a)
...

func myFunction(ptr *int) {}
```

**Nueva función**

También hay otra forma de inicializar un puntero. Podemos usar el `new` función que toma un tipo como argumento, asigna suficiente memoria para acomodar un valor de ese tipo, y devuelve un puntero a la misma.

Aquí hay un ejemplo:

```go
package main

import "fmt"

func main() {
	p := new(int)
	*p = 100

	fmt.Println("value", *p)
	fmt.Println("address", p)
}
```

`$ go run main.go
value 100
address 0xc000018030`

**Puntero a un puntero**

Aquí hay una idea interesante...¿podemos crear un puntero a un puntero? ¡La respuesta es sí! Sí, podemos.

```go
package main

import "fmt"

func main() {
	p := new(int)
	*p = 100

	p1 := &p

	fmt.Println("P value", *p, " address", p)
	fmt.Println("P1 value", *p1, " address", p)

	fmt.Println("Dereferenced value", **p1)
}
```

```
$ go run main.go
P value 100  address 0xc0000be000
P1 value 0xc0000be000  address 0xc0000be000
Dereferenced value 100
```

_Observe cómo el valor de `p1` coincide con la dirección de `p`._

Además, es importante saber que los punteros en Go no admiten la aritmética del puntero como en C o C++.

```
	p1 := p * 2 // Compiler Error: invalid operation
```

Sin embargo, podemos comparar dos punteros del mismo tipo para la igualdad usando un `==` operador.

```go
p := &a
p1 := &a

fmt.Println(p == p1)
```

**Pero ¿Por qué?**

Esto nos lleva a la pregunta del millón de dólares, ¿por qué necesitamos consejos?

Bueno, no hay una respuesta definitiva para eso, y los punteros son solo otra característica útil que nos ayuda a mutar nuestros datos de manera eficiente sin copiar una gran cantidad de datos.

Por último, agregaré que si vienes de un lenguaje sin noción de punteros, no entres en pánico y trata de formar un modelo mental de cómo funcionan los punteros.

# Estructuras

Una `struct` es un tipo definido por el usuario que contiene una colección de campos con nombre. Básicamente, se utiliza para agrupar datos relacionados para formar una sola unidad.

Si vienes de un entorno orientado a objetos, piensa en las estructuras como clases ligeras que apoyan la composición pero no la herencia.

## Definición de estructuras

Podemos definir una `struct` así:

```go
type Person struct {}
```

Usamos la palabra clave `type` para introducir un nuevo tipo, seguido del nombre y luego la palabra clave `struct` para indicar que estamos definiendo una estructura.

Ahora, añadamos algunos campos:

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```

Y, si los campos tienen el mismo tipo, también podemos agruparlos:

```go
type Person struct {
    FirstName, LastName string
    Age                 int
}
```

## Declarar e inicializar

Ahora que tenemos nuestra estructura, podemos declararla igual que otros tipos de datos:

```go
func main() {
    var p1 Person

    fmt.Println("Person 1:", p1)
}
```

```
$ go run main.go
Person 1: {  0}
```

Como podemos ver, todos los campos de estructura se inicializan con sus valores cero. Entonces el `FirstName` y `LastName` están configurados como cadena vacía `""` y `Age` se establece en 0.

También podemos inicializarla como _"estructura literal"_:

```go
func main() {
    var p1 Person

    fmt.Println("Person 1:", p1)

    var p2 = Person{FirstName: "Karan", LastName: "Pratap Singh", Age: 22}

    fmt.Println("Person 2:", p2)
}
```

```
$ go run main.go
Person 1: {  0}
Person 2: {Karan Pratap Singh 22}
```

También podemos inicializar solo un subconjunto de campos:

```go
func main() {
    var p1 Person

    fmt.Println("Person 1:", p1)

    var p2 = Person{
        FirstName: "Karan",
        LastName:  "Pratap Singh",
        Age:       22,
    }

    fmt.Println("Person 2:", p2)

    var p3 = Person{
        FirstName: "Tony",
        LastName:  "Stark",
    }

    fmt.Println("Person 3:", p3)
}
```

```
$ go run main.go
Person 1: {  0}
Person 2: {Karan Pratap Singh 22}
Person 3: {Tony Stark 0}
```

Como podemos ver, el campo de edad de la persona 3 ha tomado el valor cero por defecto.

## Inicialización sin nombre de campo

Go structs también admite la inicialización sin nombres de campo:

```go
func main() {
    var p1 Person

    fmt.Println("Person 1:", p1)

    var p2 = Person{
        FirstName: "Karan",
        LastName:  "Pratap Singh",
        Age:       22,
    }

    fmt.Println("Person 2:", p2)

    var p3 = Person{
        FirstName: "Tony",
        LastName:  "Stark",
    }

    fmt.Println("Person 3:", p3)

    var p4 = Person{"Bruce", "Wayne", 40}

    fmt.Println("Person 4:", p4)
}
```

Pero aquí está la particularidad: tendremos que proporcionar todos los valores durante la inicialización o fallará:

```
$ go run main.go
# command-line-arguments
./main.go:30:27: too few values in Person{...}
```

La forma correcta sería:

```go
var p4 = Person{"Bruce", "Wayne", 40}
fmt.Println("Person 4:", p4)
```

## Estructuras anónimas

También podemos declarar una estructura anónima:

```go
func main() {
    var p1 Person

    fmt.Println("Person 1:", p1)

    var p2 = Person{
        FirstName: "Karan",
        LastName:  "Pratap Singh",
        Age:       22,
    }

    fmt.Println("Person 2:", p2)

    var p3 = Person{
        FirstName: "Tony",
        LastName:  "Stark",
    }

    fmt.Println("Person 3:", p3)

    var p4 = Person{"Bruce", "Wayne", 40}

    fmt.Println("Person 4:", p4)

    var a = struct {
        Name string
    }{"Golang"}

    fmt.Println("Anonymous:", a)
}
```

## Acceso a campos

Limpiemos un poco nuestro ejemplo y veamos cómo podemos acceder a campos individuales:

```go
func main() {
    var p = Person{
        FirstName: "Karan",
        LastName:  "Pratap Singh",
        Age:       22,
    }

    fmt.Println("FirstName", p.FirstName)
}
```

## Punteros a estructuras

También podemos crear un puntero a las estructuras:

```go
func main() {
    var p = Person{
        FirstName: "Karan",
        LastName:  "Pratap Singh",
        Age:       22,
    }

    ptr := &p

    fmt.Println((*ptr).FirstName)
    fmt.Println(ptr.FirstName)
}
```

Ambas declaraciones son iguales ya que en Go no necesitamos desreferenciar explícitamente el puntero. También podemos usar la función incorporada `new`:

```go
func main() {
    p := new(Person)

    p.FirstName = "Karan"
    p.LastName = "Pratap Singh"
    p.Age = 22

    fmt.Println("Person", p)
}
```

```
$ go run main.go
Person &{Karan Pratap Singh 22}
```

## Comparación de estructuras

Como nota al margen, dos estructuras son iguales si todos sus campos correspondientes son iguales también:

```go
func main() {
    var p1 = Person{"a", "b", 20}
    var p2 = Person{"a", "b", 20}

    fmt.Println(p1 == p2)
}
```

```
$ go run main.go
true
```

## Campos exportados

Ahora aprendamos qué son campos exportados y no exportados en una estructura. Igual que las reglas para variables y funciones, si un campo de estructura se declara con un identificador en minúscula, no se exportará y solo será visible para el paquete en el que se define:

```go
type Person struct {
    FirstName, LastName  string
    Age                  int
    zipCode              string
}
```

El campo `zipCode` no se exportará. Además, lo mismo ocurre con la estructura `Person`, si la cambiamos de nombre a `person`, tampoco se exportará:

```go
type person struct {
    FirstName, LastName  string
    Age                  int
    zipCode              string
}
```

## Incrustación y composición

Como discutimos anteriormente, Go no necesariamente admite la herencia, pero podemos hacer algo similar con la incrustación:

```go
type Person struct {
    FirstName, LastName  string
    Age                  int
}

type SuperHero struct {
    Person
    Alias string
}
```

Nuestra nueva estructura tendrá todas las propiedades de la estructura original. Y debe comportarse igual que nuestra estructura normal:

```go
func main() {
    s := SuperHero{}

    s.FirstName = "Bruce"
    s.LastName = "Wayne"
    s.Age = 40
    s.Alias = "batman"

    fmt.Println(s)
}
```

```
$ go run main.go
{{Bruce Wayne 40} batman}
```

Sin embargo, esto generalmente no se recomienda y en la mayoría de los casos, se prefiere la composición. Entonces, en lugar de incrustar, lo definiremos como un campo normal:

```go
type Person struct {
    FirstName, LastName  string
    Age                  int
}

type SuperHero struct {
    Person Person
    Alias  string
}
```

Por lo tanto, también podemos reescribir nuestro ejemplo con composición:

```go
func main() {
    p := Person{"Bruce", "Wayne", 40}
    s := SuperHero{p, "batman"}

    fmt.Println(s)
}
```

```
$ go run main.go
{{Bruce Wayne 40} batman}
```

No hay un enfoque correcto o incorrecto aquí, pero la incrustación puede ser útil en ciertos casos.

## Etiquetas de estructura

Una etiqueta struct es solo una etiqueta que nos permite adjuntar información de metadatos al campo, que se puede utilizar para comportamiento personalizado utilizando el paquete `reflect`.

Aprendamos cómo podemos definir etiquetas de estructura:

```go
type Animal struct {
    Name    string `key:"value1"`
    Age     int    `key:"value2"`
}
```

A menudo encontrará etiquetas en los paquetes de codificación, como XML, JSON, YAML, ORM y gestión de configuración.

Aquí hay un ejemplo de etiquetas para el codificador JSON:

```go
type Animal struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
}
```

## Propiedades de las estructuras

Finalmente, discutamos las propiedades de las estructuras.

Las estructuras son tipos de valor. Cuando asignamos una variable `struct` a otra, se crea y asigna una nueva copia de la `struct`.

Del mismo modo, cuando pasamos una `struct` a otra función, la función obtiene su propia copia de la `struct`:

```go
package main

import "fmt"

type Point struct {
    X, Y float64
}

func main() {
    p1 := Point{1, 2}
    p2 := p1 // Se asigna una copia de p1 a p2

    p2.X = 2

    fmt.Println(p1) // Output: {1 2}
    fmt.Println(p2) // Output: {2 2}
}
```

La estructura vacía ocupa cero bytes de almacenamiento:

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var s struct{}
    fmt.Println(unsafe.Sizeof(s)) // Output: 0
}
```

## Métodos

Hablemos de métodos, a veces también conocidos como receptores de funciones.

Técnicamente, Go no es un lenguaje de programación orientado a objetos. No tiene clases, objetos y herencia.

Sin embargo, Go tiene tipos. Y, puedes definir **métodos** en tipos.

Un método no es más que una función con un especial _receptor_ argumento. Veamos cómo podemos declarar métodos.

```go
func (variable T) Name(params) (returnTypes) {}
```

El _receptor_ argumento tiene un nombre y un tipo. Aparece entre el `func` palabra clave y el nombre del método.

Por ejemplo, definamos un `Car` estructura.

```go
type Car struct {
	Name string
	Year int
}
```

Ahora, definamos un método como `IsLatest` lo que nos dirá si se fabricó un automóvil en los últimos 5 años.

```go
func (c Car) IsLatest() bool {
	return c.Year >= 2017
}
```

**Métodos con receptores de puntero**

Todos los ejemplos que vimos anteriormente tenían un receptor de valor.

Con un receptor de valor, el método funciona con una copia del valor que se le pasa. Por lo tanto, cualquier modificación realizada al receptor dentro de los métodos no es visible para la persona que llama.

Por ejemplo, hagamos otro método llamado `UpdateName` que actualizará el nombre de la `Car`.

```go
func (c Car) UpdateName(name string) {
	c.Name = name
}
```

Ahora, corramos esto.

```go
func main() {
	c := Car{"Tesla", 2021}

	c.UpdateName("Toyota")
	fmt.Println("Car:", c)
}
```

```
$ go run main.go
Car: {Tesla 2021}
```

Parece que el nombre no se actualizó, así que ahora cambiemos nuestro receptor al tipo de puntero e intentemos de nuevo.

```go
func (c *Car) UpdateName(name string) {
	c.Name = name
}
```

```
$ go run main.go
Car: {Toyota 2021}
```

Como se esperaba, los métodos con receptores de puntero pueden modificar el valor al que apunta el receptor. Tales modificaciones son visibles también para la persona que llama del método.

**Propiedades**

¡Veamos también algunas propiedades de los métodos!

- Go es lo suficientemente inteligente como para interpretar correctamente nuestra llamada de función y, por lo tanto, las llamadas al método del receptor de puntero son solo azúcar sintáctico proporcionado por Go para mayor comodidad.

```go
(&c).UpdateName(...)
```

- También podemos omitir la parte variable del receptor si no la estamos usando.

```go
func (Car) UpdateName(...) {}
```

- Los métodos no se limitan a estructuras, sino que también se pueden usar con tipos que no son estructuras.

```go
package main

import "fmt"

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	i := MyInt(10)

	fmt.Println(i.isGreater(5))
}
```

**¿Por qué métodos en lugar de funciones?**

Entonces la pregunta es, ¿por qué usar métodos en lugar de funciones?

Como siempre, no hay una respuesta particular para esto, y de ninguna manera uno es mejor que el otro. En cambio, deben usarse adecuadamente cuando llegue la situación.

Una cosa en la que puedo pensar en este momento es que los métodos pueden ayudarnos a evitar nombrar conflictos.

Dado que un método está vinculado a un tipo en particular, podemos tener los mismos nombres de método para múltiples receptores.

Pero al final, podría reducirse a la preferencia, como _"las llamadas de método son mucho más fáciles de leer y entender que las llamadas de función"_ o al revés.

## Arrays y Rebanadas

### Arrays

**¿Qué es una matriz?**

Una matriz es una colección de tamaño fijo de elementos del mismo tipo. Los elementos de la matriz se almacenan secuencialmente y se puede acceder a ellos utilizando su `index`.

**Declaración**

Podemos declarar una matriz de la siguiente manera:

```go
var a [n]T
```

Aquí, `n` es la longitud y `T` puede ser de cualquier tipo como entero, cadena o estructuras definidas por el usuario.

Ahora, declaremos una matriz de enteros con longitud 4 e imprímala.

```go
func main() {
	var arr [4]int

	fmt.Println(arr)
}
```

```
$ go run main.go
[0 0 0 0]
```

De forma predeterminada, todos los elementos de la matriz se inicializan con el valor cero del tipo de matriz correspondiente.

**Inicialización**

También podemos inicializar una matriz usando una matriz literal.

```go
var a [n]T = [n]T{V1, V2, ... Vn}
```

```go
func main() {
	var arr = [4]int{1, 2, 3, 4}

	fmt.Println(arr)
}
```

```
$ go run main.go
[1 2 3 4]
```

Incluso podemos hacer una declaración abreviada.

```go
arr := [4]int{1, 2, 3, 4}
```

**Acceso**

Y similar a otros idiomas, podemos acceder a los elementos utilizando el `index` como se almacenan secuencialmente.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	fmt.Println(arr[0])
}
```

```
$ go run main.go
1
```

**Iteración**

Ahora, hablemos de iteración.

Por lo tanto, hay múltiples formas de iterar sobre matrices.

El primero está usando el bucle for con el `len` función que nos da la longitud de la matriz.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr[i])
	}
}
```

```
$ go run main.go
Index: 0, Element: 1
Index: 1, Element: 2
Index: 2, Element: 3
Index: 3, Element: 4
```

Otra forma es usar el `range` palabra clave con el `for` bucle.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```

```
$ go run main.go
Index: 0, Element: 1
Index: 1, Element: 2
Index: 2, Element: 3
Index: 3, Element: 4
```

Como podemos ver, nuestro ejemplo funciona igual que antes.

Pero la palabra clave de rango es bastante versátil y se puede usar de múltiples maneras.

```go
for i, e := range arr {} // Normal usage of range

for _, e := range arr {} // Omit index with _ and use element

for i := range arr {} // Use index only

for range arr {} // Simply loop over the array
```

**Multi-dimensional**

Todas las matrices que hemos creado hasta ahora son unidimensionales. También podemos crear matrices multidimensionales en Go.

Echemos un vistazo a un ejemplo:

```go
func main() {
	arr := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```

```
$ go run main.go
Index: 0, Element: [1 2 3 4]
Index: 1, Element: [5 6 7 8]
```

También podemos dejar que el compilador infiera la longitud de la matriz utilizando `...` elipses en lugar de la longitud.

```go
func main() {
	arr := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```

```
$ go run main.go
Index: 0, Element: [1 2 3 4]
Index: 1, Element: [5 6 7 8]
```

**Propiedades**

Ahora hablemos de algunas propiedades de las matrices.

La longitud de la matriz es parte de su tipo. Entonces, la matriz `a` y `b` son tipos completamente distintos, y no podemos asignar uno al otro.

Esto también significa que no podemos cambiar el tamaño de una matriz, porque cambiar el tamaño de una matriz significaría cambiar su tipo.

```go
package main

func main() {
	var a = [4]int{1, 2, 3, 4}
	var b [2]int = a // Error, cannot use a (type [4]int) as type [2]int in assignment
}
```

Las matrices en Go son tipos de valor a diferencia de otros lenguajes como C, C++ y Java, donde las matrices son tipos de referencia.

Esto significa que cuando asignamos una matriz a una nueva variable o pasamos una matriz a una función, se copia toda la matriz.

Por lo tanto, si hacemos algún cambio en esta matriz copiada, la matriz original no se verá afectada y permanecerá sin cambios.

```go
package main

import "fmt"

func main() {
	var a = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var b = a // Copy of a is assigned to b

	b[0] = "Monday"

	fmt.Println(a) // Output: [Mon Tue Wed Thu Fri Sat Sun]
	fmt.Println(b) // Output: [Monday Tue Wed Thu Fri Sat Sun]
}
```

### Slices (Rebanadas)

Sé lo que estás pensando, las matrices son útiles pero un poco inflexibles debido a la limitación causada por su tamaño fijo.

Esto nos lleva a Slice, entonces, ¿qué es una rebanada?

Una rebanada es un segmento de una matriz. Las rodajas se basan en matrices y proporcionan más potencia, flexibilidad y conveniencia.

Una rebanada consta de tres cosas:

- Una referencia de puntero a una matriz subyacente.
- La longitud del segmento de la matriz que contiene la rebanada.
- Y, la capacidad, que es el tamaño máximo hasta el cual el segmento puede crecer.

Como `len` función, podemos determinar la capacidad de una rebanada usando el incorporado `cap` función. Aquí hay un ejemplo:

```go
package main

import "fmt"

func main() {
	a := [5]int{20, 15, 5, 30, 25}

	s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5 30], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))
}
```

**Declaración**

Veamos cómo podemos declarar una rebanada.

```go
var s []T
```

Como podemos ver, no necesitamos especificar ninguna longitud. Declaremos una porción de enteros y veamos cómo funciona.

```go
func main() {
	var s []string

	fmt.Println(s)
	fmt.Println(s == nil)
}
```

```
$ go run main.go
[]
true
```

Entonces, a diferencia de las matrices, el valor cero de una rebanada es `nil`.

Otra forma es crear una porción a partir de una matriz. Dado que una rebanada es un segmento de una matriz, podemos crear una rebanada a partir del índice `low` a `high` como sigue.

```go
a[low:high]
```

```go
func main() {
	var a = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a[0:2] // Select from 0 to 2
	s2 := a[:3]  // Select first 3
	s3 := a[2:]  // Select last 2

	fmt.Println("Array:", a)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
}
```

```
$ go run main.go
Array: [C++ Go Java TypeScript]
Slice 1: [C++ Go]
Slice 2: [C++ Go Java]
Slice 3: [Java TypeScript]
```

_Falta de índice bajo implica 0 y falta de índice alto implica la longitud de la matriz subyacente (`len(a)`)._

Lo que hay que tener en cuenta aquí es que también podemos crear una rebanada de otras rebanadas y no solo matrices.

```go
var a = []string{
	"C++",
	"Go",
	"Java",
	"TypeScript",
}
```

**Iteración**

Podemos iterar sobre un segmento de la misma manera que iterar sobre una matriz, utilizando el bucle for con cualquiera de los dos `len` función o `range` palabra clave.

**Funciones**

Así que ahora, hablemos de las funciones de corte integradas proporcionadas en Go.

**copy**

El `copy()` la función copia elementos de una rebanada a otra. Se necesitan 2 rebanadas, un destino y una fuente. También devuelve el número de elementos copiados.

```go
func copy(dst, src []T) int
```

Veamos cómo podemos usarlo.

```go
func main() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))

	e := copy(s2, s1)

	fmt.Println("Src:", s1)
	fmt.Println("Dst:", s2)
	fmt.Println("Elements:", e)
}
```

```
$ go run main.go
Src: [a b c d]
Dst: [a b c d]
Elements: 4
```

Como se esperaba, nuestros 4 elementos de la sección de origen se copiaron en la sección de destino.

**append**

Ahora, veamos cómo podemos agregar datos a nuestra rebanada usando el incorporado `append` función que añade nuevos elementos al final de un segmento dado.

Se necesita una porción y un número variable de argumentos. Luego devuelve una nueva rebanada que contiene todos los elementos.

```go
append(slice []T, elems ...T) []T
```

Probémoslo en un ejemplo agregando elementos a nuestra rebanada.

```go
func main() {
	s1 := []string{"a", "b", "c", "d"}

	s2 := append(s1, "e", "f")

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
```

```
$ go run main.go
s1: [a b c d]
s2: [a b c d e f]
```

Como podemos ver, los nuevos elementos se agregaron y se devolvió una nueva rebanada.

Pero si el segmento dado no tiene suficiente capacidad para los nuevos elementos, entonces se asigna una nueva matriz subyacente con una capacidad mayor.

Todos los elementos de la matriz subyacente de la sección existente se copian en esta nueva matriz, y luego se agregan los nuevos elementos.

**Propiedades**

Finalmente, discutamos algunas propiedades de las rebanadas.

Las rodajas son tipos de referencia, a diferencia de las matrices.

Esto significa que modificar los elementos de una rebanada modificará los elementos correspondientes en la matriz referenciada.

```go
package main

import "fmt"

func main() {
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	s := a[0:2]

	s[0] = "Sun"

	fmt.Println(a) // Output: [Sun Tue Wed Thu Fri Sat Sun]
	fmt.Println(s) // Output: [Sun Tue]
}
```

Las rodajas también se pueden usar con tipos variádicos.

```go
package main

import "fmt"

func main() {
	values := []int{1, 2, 3}
	sum := add(values...)
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

## Mapas

Entonces, Go proporciona un tipo de mapa incorporado, y aprenderemos a usarlo.

Pero, la pregunta es ¿qué son los mapas? ¿Y por qué los necesitamos?

Bueno, un mapa es una colección desordenada de pares clave-valor. Mapas claves a valores. Las claves son únicas dentro de un mapa, mientras que los valores pueden no ser.

Se utiliza para búsquedas rápidas, recuperación y eliminación de datos basados en claves. Es una de las estructuras de datos más utilizadas.

### Declaración

Comencemos con la declaración.

Un mapa se declara utilizando la siguiente sintaxis:

```go
var m map[K]V
```

Donde `K` es el tipo clave y `V` es el tipo de valor.

Por ejemplo, así es como podemos declarar un mapa de `string` claves para `int` valores.

```go
func main() {
	var m map[string]int

	fmt.Println(m)
}
```

```
$ go run main.go
nil
```

Como podemos ver, el valor cero de un mapa es `nil`.

A `nil` el mapa no tiene llaves. Además, cualquier intento de agregar claves a un `nil` map dará como resultado un error de tiempo de ejecución.

### Inicialización

Hay múltiples formas de inicializar un mapa.

#### Función make

Podemos usar el incorporado `make` función, que asigna memoria para tipos de datos referenciados e inicializa sus estructuras de datos subyacentes.

```go
func main() {
	var m = make(map[string]int)

	fmt.Println(m)
}
```

```
$ go run main.go
map[]
```

#### Mapa literal

Otra forma es usar un mapa literal.

```go
func main() {
	var m = map[string]int{
		"a": 0,
		"b": 1,
	}

	fmt.Println(m)
}
```

_Tenga en cuenta que se requiere la coma final._

```
$ go run main.go
map[a:0 b:1]
```

Como siempre, también podemos usar nuestros tipos personalizados.

```go
type User struct {
	Name string
}

func main() {
	var m = map[string]User{
		"a": User{"Peter"},
		"b": User{"Seth"},
	}

	fmt.Println(m)
}
```

¡Incluso podemos eliminar el tipo de valor y Go lo resolverá!

```go
var m = map[string]User{
	"a": {"Peter"},
	"b": {"Seth"},
}
```

```
$ go run main.go
map[a:{Peter} b:{Seth}]
```

### Añadir

Ahora, veamos cómo podemos agregar un valor a nuestro mapa.

```go
func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m["c"] = User{"Steve"}

	fmt.Println(m)
}
```

```
$ go run main.go
map[a:{Peter} b:{Seth} c:{Steve}]
```

### Recuperar

También podemos recuperar nuestros valores del mapa utilizando la clave.

```go
c := m["c"]
fmt.Println("Key c:", c)
```

```
$ go run main.go
key c: {Steve}
```

**¿Qué pasa si usamos una clave que no está presente en el mapa?**

```go
d := m["d"]
fmt.Println("Key d:", d)
```

¡Sí, lo adivinaste! obtendremos el valor cero del tipo de valor del mapa.

```
$ go run main.go
Key c: {Steve}
Key d: {}
```

### Existe

Cuando recupera el valor asignado a una clave dada, también devuelve un valor booleano adicional. La variable booleana será `true` si la clave existe, y `false` de lo contrario.

Probemos esto en un ejemplo:

```go
c, ok := m["c"]
fmt.Println("Key c:", c, "Present:", ok)

d, ok := m["d"]
fmt.Println("Key d:", d, "Present:", ok)
```

```
$ go run main.go
Key c: {Steve} Present: true
Key d: {} Present: false
```

### Actualizar

También podemos actualizar el valor de una clave simplemente reasignando una clave.

```go
m["a"] = User{"Roger"}
```

```
$ go run main.go
map[a:{Roger} b:{Seth} c:{Steve}]
```

### Eliminar

O bien, podemos eliminar la clave utilizando el incorporado `delete` función.

Así es como se ve la sintaxis:

```go
delete(m, "b")
```

El primer argumento es el mapa, y el segundo es la clave que queremos eliminar.

El `delete()` la función no devuelve ningún valor. Además, no hace nada si la clave no existe en el mapa.

```
$ go run main.go
map[a:{Roger} c:{Steve}]
```

### Iteración

Similar a matrices o cortes, podemos iterar sobre mapas con el `range` palabra clave.

```go
package main

import "fmt"

func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m["c"] = User{"Steve"}

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
```

```
$ go run main.go
Key: c, Value: {Steve}
Key: a, Value: {Peter}
Key: b, Value: {Seth}
```

Tenga en cuenta que un mapa es una colección desordenada y, por lo tanto, no se garantiza que el orden de iteración de un mapa sea el mismo cada vez que iteramos sobre él.

### Propiedades

Por último, hablemos de las propiedades del mapa.

Los mapas son tipos de referencia, lo que significa que cuando asignamos un mapa a una nueva variable, ambos se refieren a la misma estructura de datos subyacente.

Por lo tanto, los cambios realizados por una variable serán visibles para la otra.

```go
package main

import "fmt"

type User struct {
    Name string
}

func main() {
    var m1 = map[string]User{
        "a": {"Peter"},
        "b": {"Seth"},
    }

    m2 := m1
    m2["c"] = User{"Steve"}

    fmt.Println(m1) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
    fmt.Println(m2) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
}
```
