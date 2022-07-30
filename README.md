<!-- Written by itsnexn -->
<div align="center">
  <div>
    <h1>Meowfetch</h1>
    <p>minimal, fast and customizable system information program written in go.</p>
  </div>
  <div>
    <a href="https://github.com/itsnexn/meowfetch/stargazers">
		  <img alt="Stargazers" src="https://img.shields.io/github/stars/itsnexn/meowfetch?style=for-the-badge&color=a6e3a1&labelColor=302D41">
    </a>
    <a href="https://github.com/Itsnexn/meowfetch/issues">
      <img src="https://img.shields.io/github/issues/itsnexn/meowfetch?style=for-the-badge&color=f9e2af&logoColor=D9E0EE&labelColor=302D41">
    </a>
     <a href="https://github.com/Itsnexn/meowfetch/blob/master/LICENSE.txt">
      <img src="https://img.shields.io/github/license/itsnexn/meowfetch?style=for-the-badge&color=89b4fa&logoColor=D9E0EE&labelColor=302D41">
    </a>
  </div>
</div>
<br>


<img align="left" width="400" src="https://raw.githubusercontent.com/Itsnexn/meowfetch/master/assets/default.png">

## Overview
at the moment I was learning go and after seeing [meow.sh](https://github.com/catppuccin/catppuccin/blob/main/resources/meow.sh) in the [catppuccin](https://github.com/catppuccin/catppuccin)
repo I tought with myself why not rewrite it. at first i wrote the application in **C** but after seeing that code base is getting so complicated and bloated I wrote the application in golang
which also compiles into binaries and its fast :)

and thats how it all started.
<br>
<br>

## Installation
make sure you have `go` and `coreutils`
First clone the repository in your local machine:
```
git clone https://github.com/Itsnexn/meowfetch.git
```
then you can build and install application using GNU `make`:
```
make install
```

<img align="right" width="400" src="https://raw.githubusercontent.com/Itsnexn/meowfetch/master/assets/pride.png">

## Configuration
configuration file can be found `$XDG_CONFIG_HOME/meowfetch/meow.conf` or `$HOME/.meow.conf`. configuration files are written in
toml like syntax. for more information read the [wiki](https://github.com/Itsnexn/meowfetch/issues).
<br>
<br>
<br>

## Contribution
See [CONTRIBUTING.md](https://github.com/Itsnexn/meowfetch/blob/master/CONTRIBUTING.md).
