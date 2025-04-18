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
