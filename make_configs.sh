#!/bin/sh

scripts=("keeper/fill_config.py" "executor/fill_config.py")

run_script() {
    if command -v python3 &>/dev/null; then
        python3 "$1"
    elif command -v python2 &>/dev/null; then
        python2 "$1"
    elif command -v python &>/dev/null; then
        python "$1"
    else
        echo "Python not found! Please install it."
        exit 1
    fi
}

# Запуск каждого скрипта из списка
for script in "${scripts[@]}"; do
    run_script "$script"
done
