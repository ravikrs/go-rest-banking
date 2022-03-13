# go-rest-banking
https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/learn/lecture/22630239#overview

```mermaid
flowchart LR
A[Hard] -->|Text| B(Round)
B --> C{Decision}
C -->|One| D[Result 1]
C -->|Two| E[Result 2]
```
## Nodes
```mermaid
flowchart LR
id1[This is textbox]
id2(round edge textbox)
id3([stadium shaped node])
id4[[subroutine node]]
id5[(database)]
id6((circle node))
id7>asymmetric shape]
id8{rhombus}
id9{{hexagon node}}
id10[/parallelogram/]
id11[\parallelogram\]
id12[/trapezoid\]
id13[\trapezoid/]
id14["This is the (text) in the box"]
```

## Links
```mermaid
flowchart LR
Start --> Stop
a --- b
A-- This is text on link ---B
c---|This is text on link|d
C-->|This is text on link|D
e-.->f
E-. text .->F

g==>h
G== text ==>H

i --> j & k --> l

I & J --> K & L

M --x B
```

```mermaid
flowchart TD
    A[Start] --> B{Is it?}
    B -->|Yes| C[OK]
    C --> D[Rethink]
    D --> B
    B ---->|No| E[End]
```

```mermaid
flowchart TB
    c1-->a2
    subgraph one
    a1-->a2
    end
    subgraph two
    b1-->b2
    end
    subgraph three
    c1-->c2
    end
```

```mermaid
flowchart TB
    c1-->a2
    subgraph one
    a1-->a2
    end
    subgraph two
    b1-->b2
    end
    subgraph three
    c1-->c2
    end
    one --> two
    three --> two
    two --> c2
```
```mermaid
flowchart LR
  subgraph TOP
    direction TB
    subgraph B1
        direction RL
        i1 -->f1
    end
    subgraph B2
        direction BT
        i2 -->f2
    end
  end
  A --> TOP --> B
  B1 --> B2
```

go mod init github.com/ravikrs/go-rest-banking/s2  
go mod tidy 