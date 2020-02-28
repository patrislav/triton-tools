# Triton Tools

This repository contains tools for use with the [Triton](https://github.com/patrislav/triton) architecture and CPU family.

Currently it only contains a very basic assembler (trasm) and disassembler (detrasm.)

## Usage

The following command will build all the tools and place them in the `bin` directory:

```bash
make
```

To assemble a `.tras` source into a ternary program and output it to a file:

```bash
cat path/to/file.tras | bin/trasm > path/to/output.bct
```

To disassemble a ternary program into an instruction listing and print it to standard output:

```bash
cat path/to/file.bct | bin/detrasm
```
