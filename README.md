# Gone-Forward
a self-development project by & for Ade.

## Calculator
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## CalculatorProMax
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## TicTacConnect4
TicTacConnect4 is a terminal-based game based on the famous [M,N,K-game](https://en.wikipedia.org/wiki/M,n,k-game), in which two players take turns placing an assigned piece on an M x N board with the goal of getting K of their own piece in a row. The very common TicTacToe game is the 3,3,3-game and TicTacConnect4, that i have written, is merely a variant where the board is a simple 9 to 5. The first player to have 4 consecutive pieces of theirs vertically, horizontally, or diagonally wins the game. Currently there needs to two active sides to play this game, but i hope to include more interesting features in the future where players can choose to play with the computer! Below is what the game play looks like :

##### starting a Game:
<img src="https://user-images.githubusercontent.com/39970726/217633331-8d109054-baff-4f23-9e39-f60a4d652bd4.jpg" width="700" height="400">

##### Winning a Game:
<img src="https://user-images.githubusercontent.com/39970726/217635227-725ee2a0-ef91-4f0b-89f8-773bcae58bef.jpg" width="500" height="250">

## Go-routines (Concurrency in Golang)
Goroutines are functions or methods that run concurrently with other functions or methods. This can be thought of as a lightweight thread managed by the Go runtime. This was one concept that stood out while i took the Golang course and decided to include this in this self-development project as it seems to be a concept that will be used alot writing programs in this language. The app includes two functions `siteSerial` & `sitesConcurrent` that makes a call to a third function `returnType`. `returnType` takes a url, makes a http Get request with the url and returns the `content-type` of the Header. Of the two functions `siteSerial` & `sitesConcurrent`, `siteSerial` calls the third function `returnType` normally in a loop, while `sitesConcurrent` is the concurrent version of `siteSerial` using goroutines. The difference in their speed can be seen below !

##### Measurement:
<img src="https://user-images.githubusercontent.com/39970726/217636767-3c5264a3-6d1a-4149-907d-63314146b863.jpg" width="800" height="250">

## Polynomial (next application)
Polynomial application will be a simple math application that takes a polynomial, and return a solution to the polynomial. This will be developed as i continue to explore the language, and I hope to write this particular app using Linked-List !
