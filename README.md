# Aprendiendo GO

# Capitulo 1

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

# Capitulo 3

## Interfaces

En esta sección, hablemos de las interfaces.

### ¿Qué es una interfaz?

Entonces, una interfaz en Go es un **tipo abstracto** que se define usando un conjunto de firmas de método. La interfaz define el **comportamiento** para tipos similares de objetos.

_Aquí, **comportamiento** es un término clave que discutiremos en breve._

Echemos un vistazo a un ejemplo para entender esto mejor.

Uno de los mejores ejemplos de interfaces del mundo real es la toma de corriente. Imagine que necesitamos conectar diferentes dispositivos a la toma de corriente.

Intentemos implementar esto. Estos son los tipos de dispositivos que usaremos.

```go
type mobile struct {
	brand string
}

type laptop struct {
	cpu string
}

type toaster struct {
	amount int
}

type kettle struct {
	quantity string
}

type socket struct{}
```

Ahora, definamos un `Draw` método en un tipo, digamos `mobile`. Aquí simplemente imprimiremos las propiedades del tipo.

```go
func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d", m, m.brand, power)
}
```

Genial, ahora definiremos el `Plug` método en el `socket` tipo que acepta nuestro `mobile` tipo como argumento.

```go
func (socket) Plug(device mobile, power int) {
	device.Draw(power)
}
```

Intentemos _"conectar"_ el `mobile` tipo a nuestro `socket` tipo en la `main` función.

```go
package main

import "fmt"

func main() {
	m := mobile{"Apple"}

	s := socket{}
	s.Plug(m, 10)
}
```

Y si ejecutamos esto, veremos lo siguiente.

```
$ go run main.go
main.mobile -> brand: Apple, power: 10
```

Esto es interesante, pero digamos que ahora queremos conectar nuestro `laptop` tipo.

```go
package main

import "fmt"

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50) // Error: cannot use l as mobile value in argument
}
```

Como podemos ver, esto arrojará un error.

**¿Qué debemos hacer ahora? ¿Definir otro método? Como `PlugLaptop`?**

Claro, pero cada vez que agreguemos un nuevo tipo de dispositivo, también tendremos que agregar un nuevo método al tipo de socket y eso no es ideal.

Aquí es donde el `interface` entra. Esencialmente, queremos definir un **contrato** que, en el futuro, debe implementarse.

Simplemente podemos definir una interfaz como `PowerDrawer` y usarla en nuestra `Plug` función para permitir cualquier dispositivo que cumpla con los criterios, que es que el tipo debe tener un `Draw` método que coincide con la firma que requiere la interfaz.

Y de todos modos, el socket no necesita saber nada sobre nuestro dispositivo y simplemente puede llamar al `Draw` método.

Ahora intentemos implementar nuestra `PowerDrawer` interfaz. Así es como se verá.

La convención es usar **"er"** como sufijo en el nombre. Y como discutimos anteriormente, una interfaz solo debe describir el **comportamiento esperado**. Que en nuestro caso es el `Draw` método.

```go
type PowerDrawer interface {
	Draw(power int)
}
```

Ahora, necesitamos actualizar nuestro `Plug` método para aceptar un dispositivo que implementa la `PowerDrawer` interfaz como argumento.

```go
func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}
```

Y para satisfacer la interfaz, simplemente podemos agregar `Draw` métodos para todos los tipos de dispositivos.

```go
type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}
```

¡Ahora, podemos conectar todos nuestros dispositivos al socket con la ayuda de nuestra interfaz!

```go
func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)
}
```

Y funciona tal como esperábamos.

```
$ go run main.go
main.mobile -> brand: Apple, power: 10
main.laptop -> cpu: Intel i9, power: 50
main.toaster -> amount: 4, power: 30
main.kettle -> quantity: Half Empty, power: 25
```

### Pero, ¿por qué se considera esto un concepto tan poderoso?

Bueno, una interfaz puede ayudarnos a desacoplar nuestros tipos. Por ejemplo, debido a que tenemos la interfaz, no necesitamos actualizar nuestra `socket` implementación. Solo podemos definir un nuevo tipo de dispositivo con un `Draw` método.

A diferencia de otros idiomas, las Interfaces en Go se implementan **implícitamente** entonces no necesitamos algo como una `implements` palabra clave. Esto significa que un tipo satisface una interfaz automáticamente cuando tiene _"todos los métodos"_ de la interfaz.

### Interfaz Vacía

A continuación, hablemos de la interfaz vacía. Una interfaz vacía puede tomar un valor de cualquier tipo.

Así es como lo declaramos.

```go
var x interface{}
```

**Pero, ¿por qué lo necesitamos?**

Las interfaces vacías se pueden usar para manejar valores de tipos desconocidos.

Algunos ejemplos son:

- Lectura de datos heterogéneos de una API.
- Variables de un tipo desconocido, como en la `fmt.Println` función.

Para usar un valor de tipo vacío `interface{}`, podemos usar _tipo aserción_ o un _tipo interruptor_ para determinar el tipo de valor.

### Tipo de Aserción

Un _tipo aserción_ proporciona acceso al valor concreto subyacente de un valor de interfaz.

Por ejemplo:

```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
}
```

Esta declaración afirma que el valor de la interfaz tiene un tipo concreto y asigna el valor de tipo subyacente a la variable.

También podemos probar si un valor de interfaz tiene un tipo específico.

Una afirmación de tipo puede devolver dos valores:

- El primero es el valor subyacente.
- El segundo es un valor booleano que informa si la afirmación tuvo éxito.

```go
s, ok := i.(string)
fmt.Println(s, ok)
```

Esto puede ayudarnos a probar si un valor de interfaz tiene un tipo específico o no.

En cierto modo, esto es similar a cómo leemos los valores de un mapa.

Y si este no es el caso entonces, `ok` será falso y el valor será el valor cero del tipo, y no se producirá pánico.

```go
f, ok := i.(float64)
fmt.Println(f, ok)
```

Pero si la interfaz no contiene el tipo, la declaración provocará un pánico.

```go
f = i.(float64)
fmt.Println(f) // Panic!
```

```
$ go run main.go
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64
```

### Tipo Interruptor

Aquí, una instrucción `switch` se puede usar para determinar el tipo de una variable de tipo vacía `interface{}`.

```go
var t interface{}
t = "hello"

switch t := t.(type) {
case string:
	fmt.Printf("string: %s\n", t)
case bool:
	fmt.Printf("boolean: %v\n", t)
case int:
	fmt.Printf("integer: %d\n", t)
default:
	fmt.Printf("unexpected: %T\n", t)
}
```

Y si ejecutamos esto, podemos verificar que tenemos un tipo `string`.

```
$ go run main.go
string: hello
```

### Propiedades

Discutamos algunas propiedades de las interfaces.

#### Valor cero

El valor cero de una interfaz es `nil`.

```go
package main

import "fmt"

type MyInterface interface {
	Method()
}

func main() {
	var i MyInterface

	fmt.Println(i) // Output: <nil>
}
```

#### Incrustación

Podemos incrustar interfaces como estructuras. Por ejemplo:

```go
type interface1 interface {
    Method1()
}

type interface2 interface {
    Method2()
}

type interface3 interface {
    interface1
    interface2
}
```

#### Valores

Los valores de interfaz son comparables.

```go
package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyType struct{}

func (MyType) Method() {}

func main() {
	t := MyType{}
	var i MyInterface = MyType{}

	fmt.Println(t == i)
}
```

#### Valores de Interfaz

Debajo del capó, un valor de interfaz puede considerarse como una tupla que consiste en un valor y un tipo concreto.

```go
package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyType struct {
	property int
}

func (MyType) Method() {}

func main() {
	var i MyInterface

	i = MyType{10}

	fmt.Printf("(%v, %T)\n", i, i) // Output: ({10}, main.MyType)
}
```

Con eso, cubrimos interfaces en Go.

Es una característica realmente poderosa, pero recuerda _"Mayor es la interfaz, más débil es la abstracción"_ - Rob Pike.

## Errores

En este tutorial, hablemos sobre el manejo de errores.

Observe cómo dije errores y no excepciones, ya que no hay manejo de excepciones en Go.

En cambio, podemos devolver un incorporado `error` tipo que es un tipo de interfaz.

```go
type error interface {
    Error() string
}
```

Volveremos a esto en breve. Primero, tratemos de entender lo básico.

### Lo Básico

Entonces, declaremos una simple función `Divide` que, como su nombre indica, dividirá entero `a` por `b`.

```go
func Divide(a, b int) int {
	return a/b
}
```

Genial. Ahora, queremos devolver un error, digamos, para evitar la división por cero. Esto nos lleva a la construcción de errores.

### Construyendo Errores

Hay múltiples formas de hacer esto, pero veremos las dos más comunes.

### Paquete errors

El primero es usando la función `New` proporcionada por el paquete `errors`.

```go
package main

import "errors"

func main() {}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a/b, nil
}
```

Observe cómo devolvemos un error con el resultado. Y si no hay error simplemente devolvemos `nil` como es el valor cero de un error porque, después de todo, es una interfaz.

Pero, ¿cómo lo manejamos? Para eso, llamemos a la función `Divide` en nuestra función `main`.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		fmt.Println(err)
		// Do something with the error
		return
	}

	fmt.Println(result)
	// Use the result
}

func Divide(a, b int) (int, error) {...}
```

```
$ go run main.go
cannot divide by zero
```

Como puede ver, simplemente verificamos si el error es `nil` y construimos nuestra lógica en consecuencia. Esto se considera bastante idiomático en Go y verás que esto se usa mucho.

### Usando fmt.Errorf

Otra forma de construir nuestros errores es usando la función `fmt.Errorf`.

Esta función es similar a `fmt.Sprintf` y nos permite formatear nuestro error. Pero en lugar de devolver una cadena, devuelve un error.

A menudo se usa para agregar algún contexto o detalle a nuestros errores.

```go
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}

	return a/b, nil
}
```

Y debería funcionar de manera similar.

```
$ go run main.go
cannot divide 4 by zero
```

### Errores Sentinel

Otra técnica importante en Go es definir los errores esperados para que puedan verificarse explícitamente en otras partes del código. Estos a veces se conocen como errores centinela.

```go
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {...}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a/b, nil
}
```

En Go, se considera convencional prefijar la variable con `Err`. Por ejemplo, `ErrNotFound`.

Pero, ¿cuál es el punto?

Por lo tanto, esto se vuelve útil cuando necesitamos ejecutar una rama de código diferente si se encuentra un cierto tipo de error.

Por ejemplo, ahora podemos verificar explícitamente qué error ocurrió usando la función `errors.Is`.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		switch {
        case errors.Is(err, ErrDivideByZero):
            fmt.Println(err)
			// Do something with the error
        default:
            fmt.Println("no idea!")
        }

		return
	}

	fmt.Println(result)
	// Use the result
}

func Divide(a, b int) (int, error) {...}
```

```
$ go run main.go
cannot divide by zero
```

### Errores Personalizados

Esta estrategia cubre la mayoría de los casos de uso de manejo de errores. Pero a veces necesitamos funcionalidades adicionales como valores dinámicos dentro de nuestros errores.

Antes, vimos que `error` es solo una interfaz. Básicamente, cualquier cosa puede ser un error mientras implemente el método `Error()` que devuelve un mensaje de error como una cadena.

Entonces, definamos nuestra estructura personalizada `DivisionError` que contendrá un código de error y un mensaje.

```go
package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	Code int
	Msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

func main() {...}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			Code: 2000,
			Msg:  "cannot divide by zero",
		}
	}

	return a/b, nil
}
```

Aquí, usaremos la función `errors.As` en lugar de `errors.Is` para convertir el error al tipo correcto.

```go
func main() {
	result, err := Divide(4, 0)

	if err != nil {
		var divErr DivisionError

		switch {
		case errors.As(err, &divErr):
			fmt.Println(divErr)
			// Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result)
	// Use the result
}

func Divide(a, b int) (int, error) {...}
```

```
$ go run main.go
code 2000: cannot divide by zero
```

Pero ¿cuál es la diferencia entre `errors.Is` y `errors.As`?

La diferencia es que `errors.As` comprueba si el error tiene un tipo específico, a diferencia de la función `Is`, que examina si es un objeto de error particular.

También podemos usar afirmaciones de tipo, pero no se prefieren.

```go
func main() {
	result, err := Divide(4, 0)

	if e, ok := err.(DivisionError); ok {
		fmt.Println(e.Code, e.Msg) // Output: 2000 cannot divide by zero
		return
	}

	fmt.Println(result)
}
```

Por último, diré que el manejo de errores en Go es bastante diferente en comparación con el tradicional idioma `try/catch` en otros lenguajes. Pero es muy poderoso, ya que alienta al desarrollador a manejar el error de una manera explícita, lo que también mejora la legibilidad.

# Pánico y Recuperación

Antes, aprendimos que la forma idiomática de manejar condiciones anormales en un programa Go está usando errores. Si bien los errores son suficientes para la mayoría de los casos, hay algunas situaciones en las que el programa no puede continuar.

En esos casos, podemos usar el incorporado `panic` función.

## Pánico

```go
func panic(interface{})
```

El pánico es una función incorporada que detiene la ejecución normal de la corriente `goroutine`. Cuando una función llama `panic`, la ejecución normal de la función se detiene inmediatamente y el control se devuelve a la persona que llama. Esto se repite hasta que el programa sale con el mensaje de pánico y el rastro de la pila.

_Nota: Discutiremos `goroutines` más tarde en el curso._

Entonces, veamos cómo podemos usar el `panic` función.

```go
package main

func main() {
	WillPanic()
}

func WillPanic() {
	panic("Woah")
}
```

Y si corremos esto, podemos ver `panic` en acción.

```
$ go run main.go
panic: Woah

goroutine 1 [running]:
main.WillPanic(...)
        .../main.go:8
main.main()
        .../main.go:4 +0x38
exit status 2
```

Como era de esperar, nuestro programa imprimió el mensaje de pánico, seguido por el rastro de la pila, y luego se terminó.

Entonces, la pregunta es, ¿qué hacer cuando ocurre un pánico inesperado?

## Recuperar

Bueno, es posible recuperar el control de un programa de pánico utilizando el incorporado `recover` función, junto con el `defer` palabra clave.

```go
func recover() interface{}
```

Probemos un ejemplo creando un `handlePanic` función. Y luego, podemos llamarlo usando `defer`.

```go
package main

import "fmt"

func main() {
	WillPanic()
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered:", data)
}

func WillPanic() {
	defer handlePanic()

	panic("Woah")
}
```

```
$ go run main.go
Recovered: Woah
```

Como podemos ver, nuestro pánico se recuperó y ahora nuestro programa puede continuar la ejecución.

Por último, mencionaré eso `panic` y `recover` puede considerarse similar a la `try/catch` idioma en otros idiomas. Pero un factor importante es que debemos evitar el pánico y recuperarnos y usar errores cuando sea posible.

Si es así, entonces esto nos lleva a la pregunta, ¿cuándo debemos usar `panic`?

## Casos de Uso

Hay dos casos de uso válidos para `panic`:

- **Un error irrecuperable**  
  Que puede ser una situación en la que el programa no puede simplemente continuar su ejecución.
  Por ejemplo, leer un archivo de configuración que es importante para iniciar el programa, ya que no hay nada más que hacer si el archivo leído en sí falla.

- **Error del desarrollador**  
  Esta es la situación más común. Por ejemplo, desreferenciar un puntero cuando el valor es `nil` causará pánico.

## Generics en Go

En esta sección, aprenderemos sobre Generics, que es una característica muy esperada que se lanzó con Go versión 1.18.

### ¿Qué son los genéricos?

Genéricos significa tipos parametrizados. En pocas palabras, los genéricos permiten a los programadores escribir código donde el tipo se puede especificar más adelante porque el tipo no es inmediatamente relevante.

Echemos un vistazo a un ejemplo para entender esto mejor.

Para nuestro ejemplo, tenemos funciones de suma simples para diferentes tipos, tales como `int`, `float64`, y `string`. Dado que la anulación del método no está permitida en Go, generalmente tenemos que crear nuevas funciones.

```go
package main

import "fmt"

func sumInt(a, b int) int {
	return a + b
}

func sumFloat(a, b float64) float64 {
	return a + b
}

func sumString(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(sumInt(1, 2))
	fmt.Println(sumFloat(4.0, 2.0))
	fmt.Println(sumString("a", "b"))
}
```

Como podemos ver, aparte de los tipos, estas funciones son bastante similares.

Veamos cómo podemos definir una función genérica.

```go
func fnName[T constraint]() {
	...
}
```

Aquí, `T` es nuestro parámetro de tipo y `constraint` será la interfaz que permita implementar cualquier tipo de interfaz.

Lo sé, lo sé, esto es confuso. Entonces, comencemos a construir nuestro genérico `sum` función.

Aquí, lo usaremos `T` como nuestro parámetro de tipo con un vacío `interface{}` como nuestra restricción.

```go
func sum[T interface{}](a, b T) T {
	fmt.Println(a, b)
}
```

Además, comenzando con Go 1.18 podemos usar `any`, que es más o menos equivalente a la interfaz vacía.

```go
func sum[T any](a, b T) T {
	fmt.Println(a, b)
}
```

Con los parámetros de tipo, viene la necesidad de pasar argumentos de tipo, lo que puede hacer que nuestro código sea detallado.

```go
sum[int](1, 2) // explicit type argument
sum[float64](4.0, 2.0)
sum[string]("a", "b")
```

Afortunadamente, Go 1.18 viene con **tipo inferencia** lo que nos ayuda a escribir código que llama funciones genéricas sin tipos explícitos.

```go
sum(1, 2)
sum(4.0, 2.0)
sum("a", "b")
```

Ejecutemos esto y veamos si funciona.

```
$ go run main.go
1 2
4 2
a b
```

Ahora, actualicemos el `sum` función para agregar nuestras variables.

```go
func sum[T any](a, b T) T {
	return a + b
}
```

```go
fmt.Println(sum(1, 2))
fmt.Println(sum(4.0, 2.0))
fmt.Println(sum("a", "b"))
```

Pero ahora si ejecutamos esto, obtendremos un error de ese operador `+` no está definido en la restricción.

```
$ go run main.go
./main.go:6:9: invalid operation: operator + not defined on a (variable of type T constrained by any)
```

Mientras que la restricción de tipo `any` generalmente funciona, no es compatible con los operadores.

Así que definamos nuestra propia restricción personalizada utilizando una interfaz. Nuestra interfaz debe definir un conjunto de tipos que contenga `int`, `float`, y `string`.

Así es como nuestro `SumConstraint` la interfaz se ve.

```go
type SumConstraint interface {
	int | float64 | string
}

func sum[T SumConstraint](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(4.0, 2.0))
	fmt.Println(sum("a", "b"))
}
```

Y esto debería funcionar como se esperaba.

```
$ go run main.go
3
6
ab
```

También podemos usar el `constraints` paquete que define un conjunto de restricciones útiles que se utilizarán con parámetros de tipo.

```go
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}
```

Para eso, tendremos que instalar el `constraints` paquete.

```
$ go get golang.org/x/exp/constraints
go: added golang.org/x/exp v0.0.0-20220414153411-bcd21879b8fd
```

```go
import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(4.0, 2.0))
	fmt.Println(sum("a", "b"))
}
```

Aquí estamos usando el `Ordered` restricción.

```go
type Ordered interface {
	Integer | Float | ~string
}
```

`~` es un nuevo token agregado a Go y la expresión `~string` significa el conjunto de todos los tipos cuyo tipo subyacente es `string`.

Y todavía funciona como se esperaba.

```
$ go run main.go
3
6
ab
```

Los genéricos son una característica sorprendente porque permiten escribir funciones abstractas que pueden reducir drásticamente la duplicación de código en ciertos casos.

### Cuándo usar genéricos

Entonces, ¿cuándo usar genéricos? Podemos tomar los siguientes casos de uso como ejemplo:

- Funciones que operan en matrices, cortes, mapas y canales.
- Estructuras de datos de propósito general como pila o lista vinculada.
- Para reducir la duplicación de código.

Por último, agregaré que si bien los genéricos son una gran adición al lenguaje, deben usarse con moderación.

Y, se recomienda comenzar simple y solo escribir código genérico una vez que hayamos escrito código muy similar al menos 2 o 3 veces.

# Capítulo 4

## Concurrencia

En esta lección, aprenderemos sobre la concurrencia, que es una de las características más poderosas de Go.

Entonces, comencemos preguntando ¿qué es _"concurrencia"_?

### Qué es la Concurrencia

La concurrencia, por definición, es la capacidad de descomponer un programa de computadora o algoritmo en partes individuales, que se pueden ejecutar de forma independiente.

El resultado final de un programa concurrente es el mismo que el de un programa que se ha ejecutado secuencialmente.

Usando concurrencia, podemos lograr los mismos resultados en menos tiempo, aumentando así el rendimiento general y la eficiencia de nuestros programas.

Imagina a un chef solitario en una cocina pequeña. Tiene varios platos que preparar. En lugar de terminar un plato por completo antes de empezar el siguiente, el chef puede:

1. Poner a hervir agua para la pasta.
2. Mientras el agua se calienta, empieza a cortar verduras para la ensalada.
3. Luego, revisa el agua, sazona la carne y la pone a cocinar.
4. Vuelve a la ensalada, aliñándola.
5. Finalmente, cocina la pasta que ya hirvió.

El chef no está haciendo todo al mismo _instante_, pero está gestionando su tiempo y los recursos (la cocina, los utensilios) de manera eficiente para avanzar en múltiples tareas dentro de un período de tiempo. Esto es la concurrencia en esencia.

### Concurrencia vs Paralelismo

Mucha gente confunde la concurrencia con el paralelismo porque ambos implican de alguna manera ejecutar código simultáneamente, pero son dos conceptos completamente diferentes.

La concurrencia es la tarea de ejecutar y administrar múltiples cálculos al mismo tiempo, mientras que el paralelismo es la tarea de ejecutar múltiples cálculos simultáneamente.

Una simple cita de Rob Pike lo resume bastante bien:

> "La concurrencia se trata de lidiar con muchas cosas a la vez. El paralelismo se trata de hacer muchas cosas a la vez"

Pero la concurrencia en Go es más que solo sintaxis. Para aprovechar el poder de Go, primero debemos entender cómo Go se acerca a la ejecución concurrente de código. Go se basa en un modelo de concurrencia llamado CSP (Communicating Sequential Processes).

### Comunicación de Procesos Secuenciales (CSP)

[Comunicar Procesos Secuenciales](https://dl.acm.org/doi/10.1145/359576.359585) (CSP) es un modelo presentado por [Tony Hoare](https://en.wikipedia.org/wiki/Tony_Hoare) en 1978 que describe las interacciones entre procesos concurrentes. Hizo un gran avance en Informática, especialmente en el campo de la concurrencia.

Idiomas como Go y Erlang han sido altamente inspirados por el concepto de comunicar procesos secuenciales (CSP).

La concurrencia es difícil, pero CSP nos permite dar una mejor estructura a nuestro código concurrente y proporciona un modelo para pensar en la concurrencia de una manera que lo hace un poco más fácil. Aquí, los procesos son independientes y se comunican compartiendo canales entre ellos.

### Conceptos Básicos

Ahora, familiaricémonos con algunos conceptos básicos de concurrencia.

#### Carrera de Datos

Una carrera de datos ocurre cuando los procesos tienen que acceder al mismo recurso simultáneamente.

_Por ejemplo, un proceso lee mientras que otro escribe simultáneamente en el mismo recurso exacto._

#### Condiciones de Carrera

Una condición de carrera ocurre cuando el momento u orden de los eventos afecta la corrección de una pieza de código.

#### Deadlocks (Puntos Muertos)

Un punto muerto ocurre cuando todos los procesos se bloquean mientras se esperan entre sí y el programa no puede continuar.

#### Condiciones de Coffman

Hay cuatro condiciones, conocidas como las condiciones de Coffman, todas ellas deben estar satisfechas para que ocurra un punto muerto:

- **Exclusión Mutua**: Un proceso concurrente contiene al menos un recurso a la vez, lo que lo hace no compartible.

  _En el siguiente diagrama, hay una sola instancia del Recurso 1 y está en manos del Proceso 1 solamente._

- **Espera y Retención**: Un proceso concurrente contiene un recurso y está esperando un recurso adicional.

- **Sin Preferencia**: Un recurso mantenido por un proceso concurrente no puede ser quitado por el sistema. Solo puede ser liberado por el proceso que lo sostiene.

- **Espera Circular**: Un proceso está esperando el recurso en poder del segundo proceso, que está esperando el recurso en poder del tercer proceso, y así sucesivamente, hasta que el último proceso está esperando un recurso en poder del primer proceso. Por lo tanto, formando una cadena circular.

#### Livelocks (Bloqueos Vivos)

Los Livelocks son procesos que realizan activamente operaciones concurrentes, pero estas operaciones no hacen nada para avanzar en el estado del programa.

#### Starvation (Inanición)

La inanición ocurre cuando un proceso se ve privado de los recursos necesarios y no puede completar su función.

La inanición puede ocurrir debido a bloqueos o algoritmos de programación ineficientes para los procesos. Para resolver la inanición, necesitamos emplear mejores algoritmos de asignación de recursos que aseguren que cada proceso obtenga su parte justa de recursos.

## Goroutines

Piensa en una goroutine como una **función ligera e independiente** que se ejecuta al mismo tiempo que el resto de tu programa. La palabra clave `go` es lo que inicia una nueva goroutine.

Pero antes de comenzar nuestra discusión, quería compartir un importante proverbio de Go:

> "No te comuniques compartiendo memoria, comparte memoria comunicándote." - Rob Pike

### ¿Qué es una goroutine?

Una _goroutine_ es un hilo ligero de ejecución que es administrado por el tiempo de ejecución Go y esencialmente vamos a escribir código asíncrono de una manera sincrónica.

Es importante saber que no son hilos de SO reales y la función principal en sí se ejecuta como una goroutine.

Un solo hilo puede ejecutar miles de goroutines en ellos utilizando el programador de tiempo de ejecución Go que utiliza la programación cooperativa. Esto implica que si la goroutine actual está bloqueada o se ha completado, el programador moverá las otras goroutines a otro hilo OS. Por lo tanto, logramos eficiencia en la programación donde ninguna rutina está bloqueada para siempre.

Podemos convertir cualquier función en una goroutine simplemente usando la palabra clave `go`.

```go
go fn(x, y, z)
```

Antes de escribir cualquier código, es importante discutir brevemente el modelo de fork-join.

### Modelo Fork-Join

Go utiliza la idea del modelo de concurrencia de fork-join detrás de las goroutines. El modelo de fork-join implica esencialmente que un proceso secundario se divide de su proceso padre para ejecutarse simultáneamente con el proceso padre. Después de completar su ejecución, el proceso del hijo se fusiona de nuevo en el proceso del padre. El punto donde se une se llama el **punto de unión**.

Ahora, escribamos un código y creemos nuestra propia goroutine.

```go
package main

import "fmt"

func speak(arg string) {
	fmt.Println(arg)
}

func main() {
	go speak("Hello World")
}
```

Aquí la llamada de función `speak` tiene un prefijo con la palabra clave `go`. Esto le permitirá funcionar como una goroutine separada. Y eso es todo, acabamos de crear nuestra primera goroutine. ¡Es así de simple!

Genial, corramos esto:

```
$ go run main.go
```

Interesante, parece que nuestro programa no se ejecutó por completo, ya que le falta algo de salida. Esto se debe a que nuestra goroutina principal salió y no esperó a la goroutine que creamos.

¿Qué pasa si hacemos que nuestro programa espere usando la función `time.Sleep`?

```go
func main() {
	...
	time.Sleep(1 * time.Second)
}
```

```
$ go run main.go
Hello World
```

Ahí vamos, podemos ver nuestra salida completa ahora.

**Bien, así que esto funciona, pero no es ideal. Entonces, ¿cómo mejoramos esto?**

Bueno, la parte más difícil del uso de goroutines es saber cuándo se detendrán. Es importante saber que las goroutines se ejecutan en el mismo espacio de direcciones, por lo que el acceso a la memoria compartida debe estar sincronizado.

## Canales

En esta lección, aprenderemos sobre los canales.

### ¿Qué son los canales?

Bueno, simplemente definido, un canal es una tubería de comunicaciones entre goroutinas. Las cosas van en un extremo y salen en otro en el mismo orden hasta que se cierra el canal.

Como aprendimos anteriormente, los canales en Go se basan en la Comunicación de Procesos Secuenciales (CSP).

### Crear un canal

Ahora que entendemos qué son los canales, veamos cómo podemos declararlos.

```go
var ch chan T
```

Aquí, prefijamos nuestro tipo `T` (que es el tipo de datos del valor que queremos enviar y recibir) con la palabra clave `chan` que significa un canal.

Intentemos imprimir el valor de nuestro canal `ch` de tipo `string`.

```go
func main() {
    var ch chan string

    fmt.Println(ch)
}
```

```
$ go run main.go
<nil>
```

Como podemos ver, el valor cero de un canal es `nil` y si intentamos enviar datos por el canal, nuestro programa entrará en pánico.

Por lo tanto, de manera similar a los slices, podemos inicializar nuestro canal utilizando la función incorporada `make`.

```go
func main() {
    ch := make(chan string)

    fmt.Println(ch)
}
```

Y si ejecutamos esto, podemos ver que nuestro canal fue inicializado.

```
$ go run main.go
0x1400010e060
```

### Envío y recepción de datos

Ahora que tenemos una comprensión básica de los canales, implementemos nuestro ejemplo anterior usando canales para aprender cómo podemos usarlos para comunicarnos entre nuestras goroutinas.

```go
package main

import "fmt"

func speak(arg string, ch chan string) {
    ch <- arg // Send
}

func main() {
    ch := make(chan string)

    go speak("Hello World", ch)

    data := <-ch // Receive
    fmt.Println(data)
}
```

Observe cómo podemos enviar datos utilizando la sintaxis `channel <- data` y recibir datos usando la sintaxis `data := <-channel`.

```
$ go run main.go
Hello World
```

Perfecto, nuestro programa funcionó como esperábamos.

### Canales Buffereados

También tenemos canales almacenados en búfer que aceptan un número limitado de valores sin un receptor correspondiente para esos valores.

Esta _longitud del buffer_ o _capacidad_ se puede especificar usando el segundo argumento para la función `make`.

```go
func main() {
    ch := make(chan string, 2)

    go speak("Hello World", ch)
    go speak("Hi again", ch)

    data1 := <-ch
    fmt.Println(data1)

    data2 := <-ch
    fmt.Println(data2)
}
```

Debido a que este canal está almacenado en búfer, podemos enviar estos valores al canal sin una recepción concurrente correspondiente. Esto significa que _envíos_ a un canal buffereado solo bloquean cuando el buffer está lleno y _recibir_ bloquea cuando el búfer esté vacío.

Por defecto, un canal no está buffereado y tiene una capacidad de 0, por lo tanto, omitimos el segundo argumento de la función `make`.

### Canales direccionales

Cuando se utilizan canales como parámetros de función, podemos especificar si un canal está destinado a enviar o recibir solo valores. Esto aumenta la seguridad de tipo de nuestro programa, ya que por defecto un canal puede enviar y recibir valores.

En nuestro ejemplo, podemos actualizar el segundo argumento de la función `speak` de tal manera que sólo puede enviar un valor.

```go
func speak(arg string, ch chan<- string) {
    ch <- arg // Send Only
}
```

Aquí, `chan<-` solo se puede usar para enviar valores y entrará en pánico si intentamos recibir valores.

### Canales de cierre

Además, al igual que cualquier otro recurso, una vez que hayamos terminado con nuestro canal, debemos cerrarlo. Esto se puede lograr utilizando la función incorporada `close`.

Aquí, podemos pasar nuestro canal a la función `close`.

```go
func main() {
    ch := make(chan string, 2)

    go speak("Hello World", ch)
    go speak("Hi again", ch)

    data1 := <-ch
    fmt.Println(data1)

    data2 := <-ch
    fmt.Println(data2)

    close(ch)
}
```

Opcionalmente, los receptores pueden probar si un canal se ha cerrado asignando un segundo parámetro a la expresión de recepción.

```go
func main() {
    ch := make(chan string, 2)

    go speak("Hello World", ch)
    go speak("Hi again", ch)

    data1 := <-ch
    fmt.Println(data1)

    data2, ok := <-ch
    fmt.Println(data2, ok)

    close(ch)
}
```

Si `ok` es `false` entonces no hay más valores para recibir y el canal está cerrado.

_En cierto modo, esto es similar a cómo verificamos si existe una clave o no en un mapa._

### Propiedades

Por último, discutamos algunas propiedades de los canales:

- Un envío a un canal `nil` bloquea para siempre.

```go
var c chan string
c <- "Hello, World!" // Panic: all goroutines are asleep - deadlock!
```

- Recibir de un canal `nil` bloquea para siempre.

```go
var c chan string
fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!
```

- Un envío a un canal cerrado causa pánico.

```go
var c = make(chan string, 1)
c <- "Hello, World!"
close(c)
c <- "Hello, Panic!" // Panic: send on closed channel
```

- Una recepción de un canal cerrado devuelve el valor cero inmediatamente.

```go
var c = make(chan int, 2)
c <- 5
c <- 4
close(c)
for i := 0; i < 4; i++ {
    fmt.Printf("%d ", <-c) // Output: 5 4 0 0
}
```

### Recorrido de canales

También podemos usar `for` y `range` para iterar sobre los valores recibidos de un canal.

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2)

    ch <- "Hello"
    ch <- "World"

    close(ch)

    for data := range ch {
        fmt.Println(data)
    }
}
```

## Select

En este tutorial, aprenderemos sobre la declaración `select` en Go.

La instrucción `select` bloquea el código y espera múltiples operaciones de canal simultáneamente.

Una declaración `select` bloquea hasta que uno de sus casos puede ejecutarse, luego ejecuta ese caso. Elige uno al azar si varios están listos.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		one <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		two <- "Two"
	}()

	select {
	case result := <-one:
		fmt.Println("Received:", result)
	case result := <-two:
		fmt.Println("Received:", result)
	}

	close(one)
	close(two)
}
```

Similar a `switch`, `select` también tiene un caso predeterminado que se ejecuta si ningún otro caso está listo. Esto nos ayudará a enviar o recibir sin bloqueo.

```go
func main() {
	one := make(chan string)
	two := make(chan string)

	for x := 0; x < 10; x++ {
		go func() {
			time.Sleep(time.Second * 2)
			one <- "One"
		}()

		go func() {
			time.Sleep(time.Second * 1)
			two <- "Two"
		}()
	}

	for x := 0; x < 10; x++ {
		select {
		case result := <-one:
			fmt.Println("Received:", result)
		case result := <-two:
			fmt.Println("Received:", result)
		default:
			fmt.Println("Default...")
			time.Sleep(200 * time.Millisecond)
		}
	}

	close(one)
	close(two)
}
```

También es importante saber que un `select {}` vacío bloquea para siempre.

```go
func main() {
	...
	select {}

	close(one)
	close(two)
}
```

## Paquete Sync

Como aprendimos anteriormente, las goroutinas se ejecutan en el mismo espacio de direcciones, por lo que el acceso a la memoria compartida debe sincronizarse. El paquete [`sync`](https://go.dev/pkg/sync) proporciona primitivas útiles.

### WaitGroup

Un WaitGroup espera a que termine una colección de goroutinas. La goroutina principal llama a `Add` para establecer el número de goroutinas a esperar. Entonces cada una de las goroutinas se ejecuta y llama `Done` cuando termine. Al mismo tiempo, `Wait` se puede usar para bloquear hasta que todas las goroutinas hayan terminado.

### Uso

Podemos usar el `sync.WaitGroup` utilizando los siguientes métodos:

- `Add(delta int)`
  Toma un valor entero que es esencialmente el número de goroutines que tiene que esperar. Esto debe llamarse antes de ejecutar una goroutine.
- `Done()`
  Se llama dentro de la goroutina para indicar que la goroutina se ha ejecutado con éxito.
- `Wait()`
  Bloquea el programa hasta que todas las goroutinas especificadas por `Add()` han invocado `Done()` desde dentro.

### Ejemplo

Echemos un vistazo a un ejemplo.

```go
package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()
}
```

Si ejecutamos esto, podemos ver que nuestro programa se ejecuta como se esperaba.

```
$ go run main.go
working...
```

También podemos pasar el `WaitGroup` a la función directamente.

```go
func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go work(&wg)

	wg.Wait()
}
```

Pero es importante saber que un `WaitGroup` **no debe ser copiado** después del primer uso. Y si se pasa explícitamente a funciones, debe ser hecho por un _puntero_. Esto se debe a que puede afectar nuestro contador, lo que interrumpirá la lógica de nuestro programa.

Aumentemos también el número de goroutines llamando al método `Add` para esperar 4 goroutinas.

```go
func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	go work(&wg)
	go work(&wg)
	go work(&wg)
	go work(&wg)

	wg.Wait()
}
```

Y como se esperaba, todas nuestras goroutinas fueron ejecutadas.

```
$ go run main.go
working...
working...
working...
working...
```

### Mutex

Un Mutex es un bloqueo de exclusión mutua que evita que otros procesos entren en una sección crítica de datos mientras un proceso lo ocupa para evitar que ocurran condiciones de carrera.

### ¿Qué es una sección crítica?

Una sección crítica puede ser una pieza de código que no debe ejecutarse mediante múltiples subprocesos a la vez porque el código contiene recursos compartidos.

### Uso

Podemos usar `sync.Mutex` utilizando los siguientes métodos:

- `Lock()`
  Adquiere o mantiene la cerradura.
- `Unlock()`
  Libera la cerradura.
- `TryLock()`
  Intenta bloquear e informa si tuvo éxito.

### Ejemplo

Echemos un vistazo a un ejemplo, crearemos una estructura `Counter` y agregar un método `Update` que actualizará el valor interno.

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()
	fmt.Printf("Result is %d", c.value)
}
```

Corramos esto y veamos qué pasa.

```
$ go run main.go
Adding -5 to 0
Adding 10 to 0
Adding 19 to 0
Adding 25 to 0
Result is 49
```

Eso no parece exacto, parece que nuestro valor es siempre cero, pero de alguna manera obtuvimos la respuesta correcta.

Bueno, esto se debe a que, en nuestro ejemplo, múltiples goroutinas están actualizando la variable `value`. Y como debes haber adivinado, esto no es ideal.

Este es el caso de uso perfecto para Mutex. Entonces, comencemos usando `sync.Mutex` y envolvamos nuestra sección crítica entre los métodos `Lock()` y `Unlock()`.

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()
	fmt.Printf("Result is %d", c.value)
}
```

```
$ go run main.go
Adding -5 to 0
Adding 19 to -5
Adding 25 to 14
Adding 10 to 39
Result is 49
```

Parece que resolvimos nuestro problema y la salida también parece correcta.

_Nota: Similar a WaitGroup, un Mutex **no debe ser copiado** después del primer uso._

### RWMutex

Un RWMutex es un bloqueo de exclusión mutua lector/escritor. La cerradura puede ser retenida por un número arbitrario de lectores o un solo escritor.

En otras palabras, los lectores no tienen que esperarse el uno al otro. Solo tienen que esperar a que los escritores mantengan la cerradura.

`sync.RWMutex` por lo tanto, es preferible para los datos que se leen principalmente y el recurso que se ahorra en comparación con un `sync.Mutex` es el tiempo.

### Uso

Similar a `sync.Mutex`, podemos usar `sync.RWMutex` utilizando los siguientes métodos:

- `Lock()`
  Adquiere o mantiene la cerradura.
- `Unlock()`
  Libera la cerradura.
- `RLock()`
  Adquiere o mantiene el bloqueo de lectura.
- `RUnlock()`
  Libera el bloqueo de lectura.

_Observe cómo RWMutex tiene métodos adicionales `RLock` y `RUnlock` comparados con Mutex._

### Ejemplo

Agreguemos un método `GetValue` que leerá el valor del contador. También cambiaremos `sync.Mutex` a `sync.RWMutex`.

Ahora, simplemente podemos usar los métodos `RLock` y `RUnlock` para que los lectores no tengan que esperar el uno al otro.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	m     sync.RWMutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.Lock()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func (c *Counter) GetValue(wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.RLock()
	defer c.m.RUnlock()
	fmt.Println("Get value:", c.value)
	time.Sleep(400 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.GetValue(&wg)
	go c.GetValue(&wg)
	go c.GetValue(&wg)

	wg.Wait()
}
```

```
$ go run main.go
Get value: 0
Adding 10 to 0
Get value: 10
Get value: 10
```

_Nota: Ambos `sync.Mutex` y `sync.RWMutex` implementan la interfaz `sync.Locker`._

```go
type Locker interface {
    Lock()
    Unlock()
}
```

### Cond

La variable de condición `sync.Cond` se puede usar para coordinar a las goroutinas que desean compartir recursos. Cuando el estado de los recursos compartidos cambia, se puede utilizar para notificar a las goroutinas bloqueadas por un mutex.

Cada Cond tiene un bloqueo asociado (a menudo un `*Mutex` o `*RWMutex`), que debe mantenerse al cambiar la condición y al llamar al método Wait.

### ¿Por qué lo necesitamos?

Un escenario puede ser cuando un proceso está recibiendo datos, y otros procesos deben esperar a que este proceso reciba datos antes de que puedan leer los datos correctos.

Si simplemente usamos un [canal](https://karanpratapsingh.com/courses/go/channels) o mutex, solo un proceso puede esperar y leer los datos. No hay forma de notificar a otros procesos para leer los datos. Por lo tanto, podemos usar `sync.Cond` para coordinar los recursos compartidos.

### Uso

`sync.Cond` viene con los siguientes métodos:

- `NewCond(l Locker)`
  Devuelve un nuevo Cond.
- `Broadcast()`
  Despierta a todas las goroutinas esperando la condición.
- `Signal()`
  Despierta a una goroutine esperando la condición si hay alguna.
- `Wait()`
  Desbloquea atómicamente el bloqueo mutex subyacente.

### Ejemplo

Aquí hay un ejemplo que demuestra la interacción de diferentes goroutinas usando el `Cond`.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()

	fmt.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	var m sync.Mutex
	cond := sync.NewCond(&m)

	go read("Reader 1", cond)
	go read("Reader 2", cond)
	go read("Reader 3", cond)
	write("Writer", cond)

	time.Sleep(4 * time.Second)
}
```

```
$ go run main.go
Writer starts writing
Writer wakes all
Reader 2 starts reading
Reader 3 starts reading
Reader 1 starts reading
```

Como podemos ver, los lectores fueron suspendidos usando el método `Wait` hasta que el escritor usó el método `Broadcast` para despertar el proceso.

### Once

Once asegura que solo se llevará a cabo una ejecución incluso entre varias goroutinas.

### Uso

A diferencia de otros primitivos, `sync.Once` solo tiene un método único:

- `Do(f func())`
  Llama a la función `f` **solo una vez**. Si `Do` se llama varias veces, solo la primera llamada invocará la función `f`.

### Ejemplo

Esto parece bastante sencillo, tomemos un ejemplo:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
```

```
$ go run main.go
Count is 1
```

Como podemos ver, incluso cuando ejecutamos 100 goroutinas, el recuento solo se incrementó una vez.

### Pool

Pool es un grupo escalable de objetos temporales y también es seguro para la concurrencia. Cualquier valor almacenado en el grupo se puede eliminar en cualquier momento sin recibir notificación. Además, bajo carga alta, el grupo de objetos se puede expandir dinámicamente, y cuando no se usa o la concurrencia no es alta, el grupo de objetos se reducirá.

_La idea clave es la reutilización de objetos para evitar la creación y destrucción repetidas, lo que afectará el rendimiento._

### ¿Por qué lo necesitamos?

El propósito de Pool es almacenar en caché los artículos asignados pero no utilizados para su posterior reutilización, aliviando la presión sobre el recolector de basura. Es decir, facilita la creación de listas gratuitas eficientes y seguras para los hilos. Sin embargo, no es adecuado para todas las listas gratuitas.

El uso apropiado de un Pool es administrar un grupo de artículos temporales compartidos silenciosamente y potencialmente reutilizados por clientes independientes concurrentes de un paquete. Pool proporciona una forma de distribuir el costo de los gastos generales de asignación entre muchos clientes.

_Es importante tener en cuenta que Pool también tiene su costo de rendimiento. Es mucho más lento usar `sync.Pool` que la simple inicialización. Además, un Pool no debe copiarse después del primer uso._

### Uso

`sync.Pool` nos da los siguientes métodos:

- `Get()`
  Selecciona un elemento arbitrario del Pool, lo elimina del Pool y lo devuelve a la persona que llama.
- `Put(x any)`
  Agrega el artículo a la piscina.

### Ejemplo

Ahora, veamos un ejemplo.

Primero, crearemos un nuevo `sync.Pool`, donde podemos especificar opcionalmente una función para generar un valor cuando llamamos `Get`, de lo contrario devolverá un valor `nil`.

```go
package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

var pool = sync.Pool{
	New: func() any {
		fmt.Println("Creating a new person...")
		return &Person{}
	},
}

func main() {
	person := pool.Get().(*Person)
	fmt.Println("Get object from sync.Pool for the first time:", person)

	fmt.Println("Put the object back in the pool")
	pool.Put(person)

	person.Name = "Gopher"
	fmt.Println("Set object property name:", person.Name)

	fmt.Println("Get object from pool again (it's updated):", pool.Get().(*Person))
	fmt.Println("There is no object in the pool now (new one will be created):", pool.Get().(*Person))
}
```

Y si ejecutamos esto, veremos una salida interesante:

```
$ go run main.go
Creating a new person...
Get object from sync.Pool for the first time: &{}
Put the object back in the pool
Set object property name: Gopher
Get object from pool again (it's updated): &{Gopher}
Creating a new person...
There is no object in the pool now (new one will be created): &{}
```

_Observa cómo hicimos [aserción de tipo](https://karanpratapsingh.com/courses/go/interfaces#type-assertion) cuando llamamos `Get`._

Se puede ver que el `sync.Pool` es estrictamente un grupo de objetos temporales, que es adecuado para almacenar algunos objetos temporales que se compartirán entre las goroutinas.

### Map

El Map es como el estándar `map[any]any` pero es seguro para uso concurrente por múltiples goroutinas sin bloqueo o coordinación adicional. Las cargas, las tiendas y las eliminaciones se distribuyen en un tiempo constante.

### ¿Por qué lo necesitamos?

El tipo Map es _especializado_. La mayoría del código debería usar un mapa Go simple, con bloqueo o coordinación separados, para una mejor seguridad de tipo y para facilitar el mantenimiento de otros invariantes junto con el contenido del mapa.

El tipo Map está optimizado para dos casos de uso común:

- Cuando la entrada para una clave dada solo se escribe una vez, pero se lee muchas veces, como en cachés que solo crecen.
- Cuando varias goroutinas leen, escriben y sobrescriben entradas para conjuntos de claves disjuntos. En estos dos casos, el uso de un `sync.Map` puede reducir significativamente la contención de bloqueo en comparación con un mapa Go emparejado con un `Mutex` o `RWMutex` separado.

_El mapa cero está vacío y listo para usar. Un mapa no debe copiarse después del primer uso._

### Uso

`sync.Map` nos da los siguientes métodos:

- `Delete()`
  Elimina el valor de una clave.
- `Load(key any)`
  Devuelve el valor almacenado en el mapa para una clave, o nil si no hay ningún valor presente.
- `LoadAndDelete(key any)`
  Elimina el valor de una clave, devolviendo el valor anterior si lo hay. El resultado cargado informa si la clave estaba presente.
- `LoadOrStore(key, value any)`
  Devuelve el valor existente para la clave si está presente. De lo contrario, almacena y devuelve el valor dado. El resultado cargado es verdadero si el valor se cargó y falso si se almacenó.
- `Store(key, value any)`
  Establece el valor de una clave.
- `Range(f func(key, value any) bool)`
  Llama a `f` secuencialmente para cada clave y valor presente en el mapa. Si `f` devuelve falso, el rango detiene la iteración.

_Nota: El rango no corresponde necesariamente a ninguna instantánea consistente del contenido del Map._

### Ejemplo

Veamos un ejemplo. Aquí, lanzaremos un montón de goroutines que agregarán y recuperarán valores de nuestro mapa simultáneamente.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(10)
	for i := 0; i <= 4; i++ {
		go func(k int) {
			v := fmt.Sprintf("value %v", k)

			fmt.Println("Writing:", v)
			m.Store(k, v)
			wg.Done()
		}(i)
	}

	for i := 0; i <= 4; i++ {
		go func(k int) {
			v, _ := m.Load(k)
			fmt.Println("Reading: ", v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
```

Como se esperaba, nuestra operación de almacenamiento y recuperación será segura para uso concurrente.

```
$ go run main.go
Reading: <nil>
Writing: value 0
Writing: value 1
Writing: value 2
Writing: value 3
Writing: value 4
Reading: value 0
Reading: value 1
Reading: value 2
Reading: value 3
```

### Atomic

El paquete atomic proporciona primitivas de memoria atómica de bajo nivel para enteros y punteros que son útiles para implementar algoritmos de sincronización.

### Uso

El paquete `atomic` proporciona [varias funciones](https://pkg.go.dev/sync/atomic#pkg-functions) que hacen las siguientes 5 operaciones para tipos `int`, `uint`, y `uintptr`:

- Add
- Load
- Store
- Swap
- Compare and Swap

### Ejemplo

No podremos cubrir todas las funciones aquí. Entonces, echemos un vistazo a la función más utilizada como `AddInt32` para tener una idea.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func add(w *sync.WaitGroup, num *int32) {
	defer w.Done()
	atomic.AddInt32(num, 1)
}

func main() {
	var n int32 = 0
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i = i + 1 {
		go add(&wg, &n)
	}

	wg.Wait()

	fmt.Println("Result:", n)
}
```

Aquí, `atomic.AddInt32` garantiza que el resultado de `n` será 1000 ya que la ejecución de instrucciones de las operaciones atómicas no se puede interrumpir.

```
$ go run main.go
Result: 1000
```

## **Patrones Avanzados de Concurrencia**

En este tutorial, discutiremos algunos patrones de concurrencia avanzados en Go. A menudo, estos patrones se usan en combinación en el mundo real.

**Generador**

El patrón del generador se usa para generar una secuencia de valores que se usa para producir alguna salida.

En nuestro ejemplo, tenemos una función `generator` que simplemente devuelve un canal desde el que podemos leer los valores.

Esto funciona en el hecho de que _envíos_ y _recepciones_ bloquean hasta que el remitente y el receptor estén listos. Esta propiedad nos permitió esperar hasta que se solicite el siguiente valor.

```go
package main

import "fmt"

func main() {
	ch := generator()

	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Println("Value:", value)
	}
}

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}
```

Si ejecutamos esto, notaremos que podemos consumir valores que se produjeron a pedido.

```
$ go run main.go
Value: 0
Value: 1
Value: 2
Value: 3
Value: 4
```

_Este es un comportamiento similar al `yield` en JavaScript y Python._

**Fan-in**

El patrón de fan-in combina múltiples entradas en un solo canal de salida. Básicamente, multiplexamos nuestras entradas.

En nuestro ejemplo, creamos las entradas `i1` y `i2` usando la función `generateWork`. Entonces usamos nuestra [función variádica](https://karanpratapsingh.com/courses/go/functions#variadic-functions) `fanIn` para combinar valores de estas entradas a un solo canal de salida desde el que podemos consumir valores.

_Nota: no se garantizará el orden de entrada._

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	i1 := generateWork([]int{0, 2, 6, 8})
	i2 := generateWork([]int{1, 3, 5, 7})

	out := fanIn(i1, i2)

	for value := range out {
		fmt.Println("Value:", value)
	}
}

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan int) {
			for {
				value, ok := <-ch

				if !ok {
					wg.Done()
					break
				}

				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
```

```
$ go run main.go
Value: 0
Value: 1
Value: 2
Value: 6
Value: 8
Value: 3
Value: 5
Value: 7
```

**Fan-out**

El patrón de fan-out nos permite dividir esencialmente nuestro único canal de entrada en múltiples canales de salida. Este es un patrón útil para distribuir elementos de trabajo en múltiples actores uniformes.

En nuestro ejemplo, dividimos el canal de entrada en 4 canales de salida diferentes. Para un número dinámico de salidas, podemos fusionar las salidas en un canal compartido _"agregado"_ y usar `select`.

_Nota: el patrón de fan-out es diferente del pub/sub._

```go
package main

import "fmt"

func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	for range work {
		select {
		case value := <-out1:
			fmt.Println("Output 1 got:", value)
		case value := <-out2:
			fmt.Println("Output 2 got:", value)
		case value := <-out3:
			fmt.Println("Output 3 got:", value)
		case value := <-out4:
			fmt.Println("Output 4 got:", value)
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
```

Como podemos ver, nuestro trabajo se ha dividido entre múltiples goroutinas.

```
$ go run main.go
Output 1 got: 1
Output 2 got: 3
Output 4 got: 4
Output 1 got: 5
Output 3 got: 2
Output 3 got: 6
Output 3 got: 7
Output 1 got: 8
```

**Tubería**

El patrón de tubería es una serie de _etapas_ conectadas por canales, donde cada etapa es un grupo de goroutines que ejecutan la misma función.

En cada etapa, las goroutines:

- Reciben valores de canales _entrantes_ aguas arriba.
- Realizan alguna función en esos datos, generalmente produciendo nuevos valores.
- Envían valores a canales _salientes_ aguas abajo.

Cada etapa tiene cualquier número de canales entrantes y salientes, excepto la primera y la última etapa, que solo tienen canales salientes o entrantes, respectivamente. La primera etapa a veces se llama el _fuente_ o _productor_; la última etapa es el _sumidero_ o _consumidor_.

Al utilizar una tubería, separamos las preocupaciones de cada etapa, lo que proporciona numerosos beneficios tales como:

- Modificar etapas independientes entre sí.
- Mezclar y combinar cómo se combinan las etapas independientemente de modificar la etapa.

En nuestro ejemplo, hemos definido tres etapas, `filter`, `square`, y `half`.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	in := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out := filter(in) // Filter odd numbers
	out = square(out) // Square the input
	out = half(out)   // Half the input

	for value := range out {
		fmt.Println(value)
	}
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := i / 2
			out <- value
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
```

Parece que nuestra entrada fue procesada correctamente por la tubería de manera concurrente.

```
$ go run main.go
0
2
8
18
32
```

**Piscina de Trabajadores**

El grupo de trabajadores es un patrón realmente poderoso que nos permite distribuir el trabajo a través de múltiples trabajadores (goroutinas) simultáneamente.

En nuestro ejemplo, tenemos un canal `jobs` al que enviaremos nuestros trabajos y un canal `results` donde nuestros trabajadores enviarán los resultados una vez que hayan terminado de hacer el trabajo.

Después de eso, podemos lanzar a nuestros trabajadores simultáneamente y simplemente recibir los resultados del canal `results`.

_Idealmente, `totalWorkers` debe configurarse para `runtime.NumCPU()` lo que nos da el número de CPU lógicas utilizables por el proceso actual._

```go
package main

import (
	"fmt"
	"sync"
)

const totalJobs = 4
const totalWorkers = 2

func main() {
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	for w := 1; w <= totalWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Receive results
	for a := 1; a <= totalJobs; a++ {
		<-results
	}

	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)

		go func(job int) {
			defer wg.Done()

			fmt.Printf("Worker %d started job %d\n", id, job)

			// Do work and send result
			result := job * 2
			results <- result

			fmt.Printf("Worker %d finished job %d\n", id, job)
		}(j)
	}

	wg.Wait()
}
```

Como se esperaba, nuestros trabajos se distribuyeron entre nuestros trabajadores.

```
$ go run main.go
Worker 2 started job 4
Worker 2 started job 1
Worker 1 started job 3
Worker 2 started job 2
Worker 2 finished job 1
Worker 1 finished job 3
Worker 2 finished job 2
Worker 2 finished job 4
```

**Colas**

El patrón de cola nos permite procesar `n` número de artículos a la vez.

En nuestro ejemplo, utilizamos un canal buffer para simular un comportamiento de cola. Simplemente enviamos una [estructura vacía](https://karanpratapsingh.com/courses/go/structs#properties) a nuestro canal `queue` y esperamos a que sea liberado por el proceso anterior para que podamos continuar.

Esto es porque _enviar_ a un canal tamponado bloquea solo cuando el buffer está lleno y _recibir_ bloquea cuando el búfer está vacío.

Aquí, tenemos un trabajo total de 10 artículos y tenemos un límite de 2. Esto significa que podemos procesar 2 artículos a la vez.

_Observe cómo nuestro canal `queue` es de tipo `struct{}` ya que una estructura vacía ocupa cero bytes de almacenamiento._

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

const limit = 2
const work = 10

func main() {
	var wg sync.WaitGroup

	fmt.Println("Queue limit:", limit)
	queue := make(chan struct{}, limit)

	wg.Add(work)

	for w := 1; w <= work; w++ {
		process(w, queue, &wg)
	}

	wg.Wait()

	close(queue)
	fmt.Println("Work complete")
}

func process(work int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		fmt.Println("Processed:", work)

		<-queue
	}()
}
```

Si ejecutamos esto, notaremos que se detiene brevemente cuando cada 2º elemento (que es nuestro límite) se procesa a medida que nuestra cola espera ser eliminada.

```
$ go run main.go
Queue limit: 2
Processed: 1
Processed: 2
Processed: 4
Processed: 3
Processed: 5
Processed: 6
Processed: 8
Processed: 7
Processed: 9
Processed: 10
Work complete
```

**Patrones adicionales**

Algunos patrones adicionales que podrían ser útiles saber:

- Canal tee
- Canal puente
- Canal de búfer de anillo
- Paralelismo limitado

## **Contexto en Go**

En los programas concurrentes, a menudo es necesario adelantarse a las operaciones debido a tiempos de espera, cancelaciones o fallas de otra parte del sistema.

El paquete `context` facilita la transferencia de valores de solicitudes, señales de cancelación y plazos a través de los límites de la API a todas las goroutines involucradas en el manejo de una solicitud.

### **Tipos**

Veamos algunos tipos centrales del paquete `context`.

### **Context**

El tipo `Context` es una `interface` que se define de la siguiente manera:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

El tipo `Context` tiene los siguientes métodos:

- `Done() <-chan struct{}`
  Devuelve un canal que se cierra cuando se cancela el contexto o se agota el tiempo. Done puede devolver `nil` si el contexto nunca puede ser cancelado.
- `Deadline() (deadline time.Time, ok bool)`
  Devuelve la hora en que el contexto se cancelará o se cronometrará. Deadline devuelve `ok=false` cuando no se establece un plazo.
- `Err() error`
  Devuelve un error que explica por qué se cerró el canal Done. Si Done aún no está cerrado, regresa `nil`.
- `Value(key any) any`
  Devuelve el valor asociado con la clave o `nil` si no hay ninguno.

### **CancelFunc**

Una `CancelFunc` le dice a una operación que abandone su trabajo y no espera a que el trabajo se detenga. Si es llamado por múltiples goroutines simultáneamente, después de la primera llamada, las llamadas posteriores a un `CancelFunc` no hacen nada.

```go
type CancelFunc func()
```

### **Uso**

Revisemos las funciones que están expuestas por el paquete `context`:

### **Background**

`Background()` devuelve un `Context` no-nil y vacío. Nunca se cancela, no tiene valores y no tiene fecha límite.

_Por lo general, se utiliza por la función principal, la inicialización y las pruebas, y como el contexto de nivel superior para las solicitudes entrantes._

```go
func Background() Context
```

### **TODO**

Similar a la función `Background`, la función `TODO` también devuelve un `Context` no-nil y vacío.

Sin embargo, solo debe usarse cuando no estamos seguros de qué contexto usar o si la función no se ha actualizado para recibir un contexto. Esto significa que planeamos agregar contexto a la función en el futuro.

```go
func TODO() Context
```

### **WithValue**

Esta función toma un contexto y devuelve un contexto derivado donde el valor `val` está asociado con `key` y fluye a través del árbol de contexto.

Esto significa que una vez que obtienes un contexto con valor, cualquier contexto que se derive de este obtiene este valor.

_No se recomienda pasar parámetros críticos usando valores de contexto. En cambio, las funciones deben aceptar esos valores en la firma haciéndola explícita._

```go
func WithValue(parent Context, key, val any) Context
```

#### **Ejemplo**

Veamos un ejemplo simple de cómo agregar un par clave-valor al contexto:

```go
package main

import (
    "context"
    "fmt"
)

func main() {
    processID := "abc-xyz"

    ctx := context.Background()
    ctx = context.WithValue(ctx, "processID", processID)

    ProcessRequest(ctx)
}

func ProcessRequest(ctx context.Context) {
    value := ctx.Value("processID")
    fmt.Printf("Processing ID: %v", value)
}
```

Al ejecutar esto, veremos el `processID` siendo pasado a través de nuestro contexto:

```
$ go run main.go
Processing ID: abc-xyz
```

### **WithCancel**

Esta función crea un nuevo contexto a partir del contexto principal y devuelve el contexto derivado y la función de cancelación. El padre puede ser un `context.Background` o un contexto que se pasó a la función.

Cancelar este contexto libera recursos asociados con él, por lo que el código debe llamar a cancel tan pronto como se completen las operaciones que se ejecutan en este contexto.

_No se recomienda pasar la función `cancel`, ya que puede conducir a un comportamiento inesperado._

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

### **WithDeadline**

Esta función devuelve un contexto derivado de su padre que se cancela cuando la fecha límite excede o se llama a la función de cancelación.

Por ejemplo, podemos crear un contexto que se cancelará automáticamente en un momento determinado en el futuro y lo pasará a funciones hijas. Cuando ese contexto se cancela debido a que la fecha límite se agota, todas las funciones que obtuvieron el contexto se notifican para detener el trabajo y regresar.

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

### **WithTimeout**

Esta función es solo un envoltorio alrededor de la función `WithDeadline` con el tiempo de espera agregado.

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}
```

### **Ejemplo Completo**

Veamos un ejemplo para consolidar nuestra comprensión del contexto.

En el siguiente ejemplo, tenemos un servidor HTTP simple que maneja una solicitud:

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Handler started")
    context := req.Context()

    select {
    // Simulating some work by the server, waits 5 seconds and then responds.
    case <-time.After(5 * time.Second):
        fmt.Fprintf(w, "Response from the server")

    // Handling request cancellation
    case <-context.Done():
        err := context.Err()
        fmt.Println("Error:", err)
    }

    fmt.Println("Handler complete")
}

func main() {
    http.HandleFunc("/request", handleRequest)

    fmt.Println("Server is running...")
    http.ListenAndServe(":4000", nil)
}
```

### **Prueba del Ejemplo**

Abramos dos terminales. En la primera terminal ejecutaremos nuestro ejemplo:

```
$ go run main.go
Server is running...
Handler started
Handler complete
```

En el segundo terminal, simplemente haremos una solicitud a nuestro servidor. Si esperamos 5 segundos, obtenemos una respuesta:

```
$ curl localhost:4000/request
Response from the server
```

Ahora, veamos qué sucede si cancelamos la solicitud antes de que se complete:

_Nota: podemos usar `ctrl + c` para cancelar la solicitud a mitad de camino._

```
$ curl localhost:4000/request
^C
```

Y como podemos ver, podemos detectar la cancelación de la solicitud gracias al contexto:

```
$ go run main.go
Server is running...
Handler started
Error: context canceled
Handler complete
```

Este ejemplo muestra cómo el paquete `context` puede ser inmensamente útil para manejar cancelaciones y tiempos de espera. Podemos usarlo para cancelar cualquier trabajo intensivo en recursos si ya no es necesario, ha excedido el plazo o un tiempo de espera.
