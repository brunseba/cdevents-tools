# License

## MIT License

**Copyright (c) 2025 Sebastien Brun**

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## About This License

The MIT License is a permissive open-source license that allows:

- ✅ **Commercial use** - You can use this software for commercial purposes
- ✅ **Modification** - You can modify the software
- ✅ **Distribution** - You can distribute the software
- ✅ **Private use** - You can use the software privately
- ✅ **Sublicensing** - You can sublicense the software

### Requirements

- **License and copyright notice** - Include the original license and copyright notice

### Limitations

- ❌ **No warranty** - The software is provided "as is"
- ❌ **No liability** - Authors are not liable for damages
- ❌ **No patent rights** - This license does not grant patent rights

## Third-Party Licenses

This project uses several third-party libraries and dependencies. Here are the key ones:

### Go Dependencies

- **CDEvents SDK** - Apache License 2.0
- **CloudEvents SDK** - Apache License 2.0
- **Cobra CLI** - Apache License 2.0
- **Viper** - MIT License

### Development Tools

- **golangci-lint** - GPL-3.0 License
- **gocyclo** - BSD 3-Clause License
- **Docker** - Apache License 2.0

For a complete list of dependencies and their licenses, please check the `go.mod` file and run:

```bash
go mod download
go list -m -json all | jq -r '.Path + " " + .Version'
```

## Contributing

By contributing to this project, you agree that your contributions will be licensed under the same MIT License.

## Questions?

If you have any questions about the license or usage rights, please:

1. Check the [LICENSE](../LICENSE) file in the root directory
2. Open an issue on [GitHub](https://github.com/brunseba/cdevents-tools/issues)
3. Contact the maintainers

---

*This license information is provided for convenience. The canonical license is the [LICENSE](../LICENSE) file in the root directory of this repository.*
