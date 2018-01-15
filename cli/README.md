Example for a Command Line Interface (CLI) with [Cobra](https://github.com/spf13/cobra)

# Examples
## "cli" with command "list"

    # list all files and directories one directory above
    $ cli list ..

    # list maximal two files and directories in the current directory
    $ cli list -m 2

## "cli" with command "bash-completion"    

    # generate a bash completion file: bash-completion.sh
    $ cli bash-completion 

    # generate a bash completion file: my-bash-completion.sh
    $ cli bash-completion my-bash-completion.sh
