# Quick start

If this directory is opened as a workspace the PYTHONPATH variable needs to be
set correctly for Pytest to function appropriately via the terminal (the `.env`
file seems to work for Python functions outside the terminal). This is because
VScode terminal functions
[don't allow the necessary setting in multi-root workspace `settings.json` files](https://stackoverflow.com/questions/68900386/set-integrated-terminal-name-in-vscode-with-settings-json?noredirect=1).

To work around this issue, you need to manually run the `export_env.sh` script
to set PYTHONPATH for you when you open a terminal.
