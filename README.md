# go-on-termux

## Termux installation using apk link on GitHub
https://github.com/termux/termux-app/releases/tag/v0.118.1

Running Go code on Termux is possible by installing the Go programming language on Termux and then compiling and running your Go code. Here's a step-by-step guide:

### 1. Install Go on Termux
First, you need to install Go in Termux. You can do this by running the following commands:

```sh
pkg update
pkg upgrade
pkg install golang
```

### 2. Verify the Installation
After installation, you can verify that Go is installed by checking the version:

```sh
go version
```

This should output the installed Go version.

### 3. Set Up the Go Workspace
Next, set up your Go workspace. By default, Go uses the `~/go` directory for your workspace. You can either use this or specify a custom workspace.

```sh
mkdir -p ~/go/{bin,src,pkg}
```

### 4. Write Your Go Code
Create a Go file. For example:

```sh
cd ~/go/src
nano hello.go
```

In `hello.go`, write your Go code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Termux!")
}
```

Save and exit the file.

### 5. Run the Go Code
Compile and run your Go program:

```sh
go run hello.go
```

This should output:

```sh
Hello, Termux!
```

### 6. (Optional) Build the Go Program
If you want to build the program into a binary, you can do so with:

```sh
go build hello.go
```

Then run the compiled binary:

```sh
./hello
```

This allows you to run the program without the Go runtime. 

You're now set up to write, compile, and run Go code on Termux!

You can create a simple Go server and set up a SQLite database on Termux. Here’s a step-by-step guide:

### 1. Set Up Termux

Ensure that your Termux environment is up to date:

```sh
pkg update
pkg upgrade
termux-setup-storage
```

### 2. Install Go and SQLite

You need to install both Go and SQLite:

```sh
pkg install golang
pkg install sqlite
```

### 3. Create a Simple Go Server

Navigate to your Go workspace and create a directory for your project: (make sure you create diretory in termux home directory i.e. ~/. other than that will create error in go mode tidy step)

```sh
mkdir -p ~/go/src/myapp
cd ~/go/src/myapp
```

Create a file named `main.go`:

```sh
nano main.go
```

Add the following code to create a simple HTTP server:

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Initialize SQLite database
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a simple table
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    // Insert a user
    _, err = db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
    if err != nil {
        log.Fatal(err)
    }

    // Simple HTTP handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT name FROM users")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        for rows.Next() {
            var name string
            rows.Scan(&name)
            fmt.Fprintf(w, "User: %s\n", name)
        }
    })

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 4. Install SQLite Driver for Go

You need to install the SQLite driver for Go. Use `go get` to install it:

```sh
go get github.com/mattn/go-sqlite3
```

### 5. Run the Go Server

You can now run your Go server:

```sh
go run main.go
```

The server will start, and you can access it by navigating to `http://localhost:8080` in a browser or using `curl`:

```sh
curl http://localhost:8080
```

### 6. Manage the Database

You can also interact with the SQLite database directly using the SQLite command-line tool in Termux:

```sh
sqlite3 test.db
```

Once inside the SQLite shell, you can run SQL commands to interact with your `users` table:

```sql
SELECT * FROM users;
```

### 7. (Optional) Build and Run as a Binary

If you want to build the server into a binary and run it:

```sh
go build main.go
./main
```

Now your Go server and SQLite database are set up and running on Termux!

To install Git on Termux, follow these simple steps:

### 1. Update Termux Packages
First, ensure that your package list is up to date:

```sh
pkg update
pkg upgrade
```

### 2. Install Git
Next, install Git by running the following command:

```sh
pkg install git
```

### 3. Verify the Installation
After installation, you can verify that Git was installed correctly by checking the version:

```sh
git --version
```

This should display the installed Git version.

### 4. (Optional) Set Up Git Configuration
You may want to configure Git with your name and email address:

```sh
git config --global user.name "Your Name"
git config --global user.email "youremail@example.com"
```

Git is now installed and ready to use in Termux! You can clone repositories, commit changes, and push them to remote repositories as needed.

You can create, clone, and commit to a GitHub repository using Termux by following these steps:

### 1. Install Git and Set Up SSH (Optional)
If you haven't already installed Git, you can do so with:

```sh
pkg update
pkg install git
```

Optionally, you can set up SSH keys to avoid entering your password every time you push to GitHub.

```sh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

Add your SSH key to the GitHub account by copying the public key:

```sh
cat ~/.ssh/id_rsa.pub
```

Go to GitHub, navigate to **Settings > SSH and GPG keys**, and add the new SSH key.

### 2. Configure Git
Configure Git with your username and email:

```sh
git config --global user.name "Your Name"
git config --global user.email "youremail@example.com"
```

### 3. Create a New GitHub Repository

#### Option 1: Using the GitHub Web Interface

1. Go to [GitHub](https://github.com) and sign in.
2. Click on the "+" icon in the top-right corner and select **New repository**.
3. Fill in the repository name, description, and other settings, then click **Create repository**.
4. Copy the repository URL (either SSH or HTTPS).

#### Option 2: Using GitHub CLI

You can also use GitHub CLI to create a repository directly from Termux, but it requires installation and authentication:

```sh
pkg install gh
gh auth login
gh repo create <repo-name> --public --source=.
```

### 4. Clone a GitHub Repository

To clone the repository to your Termux environment, use the `git clone` command:

```sh
git clone <repository-url>
```

Replace `<repository-url>` with the SSH or HTTPS URL of your repository.

### 5. Add Files to the Repository

Navigate to the cloned repository:

```sh
cd <repo-name>
```

Add a new file or modify existing files. For example, create a `README.md`:

```sh
echo "# My Repository" > README.md
```

### 6. Commit Changes

Add the files to the staging area:

```sh
git add .
```

Commit the changes with a message:

```sh
git commit -m "Initial commit"
```

### 7. Push Changes to GitHub

Push the committed changes to GitHub:

```sh
git push origin main
```

If you're using a different branch name (e.g., `master`), replace `main` with your branch name.

### 8. Verify on GitHub
Go to your GitHub repository in your browser to verify that your files and commits have been successfully pushed.

Now you've successfully created, cloned, committed, and pushed changes to a GitHub repository using Termux!

The error "Key is invalid. You must supply a key in OpenSSH public key format" suggests that the key you are trying to add to GitHub is not in the correct format. Here’s how to ensure you generate and add the correct type of SSH key:

### 1. **Generate a New SSH Key in OpenSSH Format**

To generate a new SSH key in the correct OpenSSH format, use the following command:

```sh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

This command does the following:
- `-t rsa`: Specifies the type of key to create, in this case, RSA.
- `-b 4096`: Specifies the number of bits in the key (4096 is a good standard).
- `-C "your_email@example.com"`: Adds a label to the key, often your email address.

Press `Enter` to accept the default file location. You can also set a passphrase for additional security, or press `Enter` if you don't want to use a passphrase.

### 2. **Ensure the Key is in the Correct Format**

The correct OpenSSH public key format looks something like this:

```plaintext
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC3... user@example.com
```

To check the format of your key:

```sh
cat ~/.ssh/id_rsa.pub
```

The output should start with `ssh-rsa` (for RSA keys) and have a long string of characters, followed by your email or a comment at the end.

### 3. **Copy the SSH Key**

Carefully copy the entire output of the `id_rsa.pub` file:

```sh
cat ~/.ssh/id_rsa.pub | termux-clipboard-set
```

This command will copy the key to your clipboard.

### 4. **Add the SSH Key to GitHub**

1. Go to [GitHub SSH settings](https://github.com/settings/keys).
2. Click **New SSH key**.
3. Paste the copied key into the "Key" field.
4. Give it a recognizable title (like "Termux SSH Key").
5. Click **Add SSH key**.

### 5. **Test the SSH Connection**

After adding the key, test the connection:

```sh
ssh -T git@github.com
```

You should see a message like:

```sh
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```

### 6. **Retry Git Operations**

Now that your key is correctly added to GitHub, you should be able to clone, push, or pull without issues:

```sh
git clone git@github.com:username/repo.git
```

This should resolve the issue with the key format and allow you to authenticate with GitHub using SSH.

