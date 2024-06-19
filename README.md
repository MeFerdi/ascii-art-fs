
# Ascii-art-fs

This project is a Go implementation of ASCII art using the file system API. It allows users to generate ASCII art for a given string using different banner styles.


## Features

- Supports different banner styles
- Generates ASCII art for a given string
- Uses the Go file system API

## Run Locally

Clone the project

```bash
  git clone https://learn.zone01kisumu.ke/git/ferdyodhiambo/ascii-art-fs.git
```

Go to the project directory

```bash
  cd ascii-art-fs
```


```bash
  go run . <string> [bannerfile]
```


## Running Tests

To run tests, run the following command

```bash
  cd ascii
```
```bash
  go test -v
```


## Usage/Examples
```bash
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $

$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $

```


## Contributing



Step 1: **Identify an Issue**

Go to the issue tracker for your ASCII art project on GitHub:
   **https://github.com/MeFerdi/ascii-art-fs/issues**

Browse through the open issues and identify one that you are interested in fixing.

Step 2: **Create a Fork**

Fork the ASCII art project repository on GitHub:**https://github.com/MeFerdi/ascii-art-fs/fork**

Clone your forked repository to your local machine:

```bash
 git clone https://github.com/<your_github_account>/ascii-art-fs.git
```
Step 3: **Create a Branch**

    Create a new branch for your changes:

    git branch <branch_name>

Step 4: **Fix the Issue**

    Make the necessary changes to the code to fix the issue.
    Write tests to ensure that your changes work as expected.

Step 5: **Run the Tests**

    Run the tests to ensure that your changes do not break anything:
```bash
    go run . "hello" standard | cat -e
```
Step 6: **Commit and Push**

    Commit your changes:

    git add .
    git commit -m "Fixed issue #<issue_number>"

Push your changes to your fork:
```bash
git push origin <branch_name>
```

Step 7: **Create a Pull Request**

1.Go to your fork on GitHub and click on the Pull requests tab.

2.Click on the New pull request button.

3.Select your branch as the compare branch.

4.Review the changes and ensure that they are correct.

5.Click Create pull request to submit your changes for review.

## Team
> - Dennis Mbuya
> - Brian Bantu
> - Ferdynand Odhiambo

## Credit

https://www.zone01kisumu.ke
