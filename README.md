# Gone-Forward
a self-development project by & for Ade.

**Intro:** This is solely intended to be a self-development project, after acquiring a new skill (Golang). With a foundational background in C++, Golang was fairly easy for me to pick up because it is heavily influenced by C and other C-style languages. They both have a lot in common but are not exactly the same, and the purpose of this self development project is to get up to speed with the differences to see how Golang works exactly. 

A number of the applications here i have written before in C++. Especially TicTacConnect4 (I think this is the 3rd variation of the [M,N,K-game](https://en.wikipedia.org/wiki/M,n,k-game) i must have written so far haha). This was a good opportunity to refresh, learn & upskill.

To run the applications here, you can just follow the instructions at: https://go.dev/doc/tutorial/getting-started to set up your IDE, clone this repo and it should be fairly straightforward after that. Thanks!

## Calculator
The Calculator app is a very simple app initially written to get up to speed with the programming language while i took the course on LinkedIn Learning.

![Image 2-9-23 at 10 32 AM](https://user-images.githubusercontent.com/39970726/217878034-1dba473b-1b48-468c-80ca-2f41871b25ab.jpg)

## CalculatorProMax
CalculatorProMax is an extension of the Calculator app above with more features included in it. Golang doesn't exactly have the keyword class and twhis was written to get up to speed on how Golang allows the construction of a class as with other programming languages. Below is an image of how the app works:

<img src="https://user-images.githubusercontent.com/39970726/217878707-6e221eb1-b761-4f11-9b44-f45c26796d85.jpg" width="800" height="200">


## TicTacConnect4
TicTacConnect4 is a terminal-based game based on the famous [M,N,K-game](https://en.wikipedia.org/wiki/M,n,k-game), in which two players take turns placing an assigned piece on an M x N board with the goal of getting K of their own piece in a row. The very common [TicTacToe game](https://en.wikipedia.org/wiki/Tic-tac-toe) is the 3,3,3-game and TicTacConnect4, that i have written, is merely a variant where the board is a simple 9 to 5. The first player to have 4 consecutive pieces of theirs vertically, horizontally, or diagonally wins the game. Currently there needs to two active sides to play this game, but i hope to include more interesting features in the future where players can choose to play with the computer! Below is what the game play looks like :

##### starting a Game:
<img src="https://user-images.githubusercontent.com/39970726/217633331-8d109054-baff-4f23-9e39-f60a4d652bd4.jpg" width="700" height="400">

##### Winning a Game:
<img src="https://user-images.githubusercontent.com/39970726/217635227-725ee2a0-ef91-4f0b-89f8-773bcae58bef.jpg" width="500" height="250">

## Go-routines (Concurrency in Golang)
Goroutines are functions or methods that run concurrently with other functions or methods. This can be thought of as a lightweight thread managed by the Go runtime. This was one concept that stood out while i took the Golang course and decided to include this in this self-development project as it seems to be a concept that will be used alot writing programs in this language. The app includes two functions `siteSerial` & `sitesConcurrent` that makes a call to a third function `returnType`. `returnType` takes a url, makes a http Get request with the url and returns the `content-type` of the Header. Of the two functions `siteSerial` & `sitesConcurrent`, `siteSerial` calls the third function `returnType` normally in a loop, while `sitesConcurrent` is the concurrent version of `siteSerial` using goroutines. The difference in their speed can be seen below !

##### Measurement:
<img src="https://user-images.githubusercontent.com/39970726/217636767-3c5264a3-6d1a-4149-907d-63314146b863.jpg" width="800" height="250">

## AWS-SDK-GO
I have also explored the AWS Software Development Kit for Golang (check out my learning [repository](https://github.com/Ademiluyi/ade-learning-go/tree/main/s3-go-aws)). This simple application Creates an Amazon S3 service client and makes a call to retrive an image i have stored in an S3 bucket of a free AWS account. Should be able to explore more in the future.

## Polynomial (next application)
Polynomial application will be a simple math application that takes a polynomial, and return a solution to the polynomial. This will be developed as i continue to explore the language, and I hope to write this particular app using Linked-List !
