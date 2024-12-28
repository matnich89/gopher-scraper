# Email-Scraper

A high-performance web crawler and email harvesting tool written in Go, designed for security professionals conducting authorised penetration tests and security assessments.

## Motivation

While conducting research into penetration testing ( a field I have recently becoming interested in) , I noticed that many existing OSINT and reconnaissance tools are written in Python or PHP. While these tools are excellent, I wanted to create something that could leverage Go's superior performance characteristics, particularly for CPU-intensive tasks like crawling and regex matching at scale.

This tool represents an effort to modernise the pentesting toolkit by providing a fast, efficient alternative that can handle large-scale email harvesting tasks with minimal resource usage.

## Features

- Fast, concurrent web crawling (Currently not concurrent TODO)
- Efficient email pattern matching using regular expressions
- Configurable crawl depth and URL limits
- Support for both relative and absolute URLs
- Automatic deduplication of discovered email addresses
- Real-time logging of discovered emails

## Usage

```bash
./email-scraper -url https://target-domain.com -max 100
```

### Flags
- `-url`: Target URL to begin crawling (required)
- `-max`: Maximum number of URLs to crawl (default: 100)

## Intended Use Cases

This tool is designed for security professionals conducting authorized security assessments to:
- Test an organisation's resilience against social engineering attacks
- Validate email security policies
- Assess the exposure of corporate email addresses
- Identify potential phishing targets for awareness training

## Legal Disclaimer

This tool is provided for legitimate security testing purposes only. The author assumes no liability and is not responsible for any misuse or damage caused by this tool.

**You are responsible for:**
- Obtaining proper authorisation before testing any systems
- Complying with all applicable laws and regulations
- Using this tool only on systems you are explicitly authorised to test
- Any consequences of misusing this tool

## Technical Details

The tool is built using several key Go packages:
- `net/http`: For efficient HTTP requests
- `goquery`: HTML parsing and traversal
- Custom concurrent crawler implementation
- Regex-based email pattern matching

## Performance

Written in Go, this tool offers several performance advantages over similar Python/PHP implementations:
- Concurrent crawling with efficient resource usage
- Minimal memory footprint
- Fast regex pattern matching
- Efficient string manipulation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Warning

This tool should only be used for legitimate penetration testing and security assessment purposes where explicit permission has been granted. Unauthorised use of this tool may violate applicable laws and regulations.