# PlaNet

<div align="left">
    <br />
    <img src="src\assets\PlaNet.png" alt="Logo" width="80">
    <br />
    <br />
</div>

PlaNet is a convenient system tray application for Windows that enables users to quickly set and reset their DNS settings to predefined configurations.

## Features

- **Set DNS:** Change your current DNS to the Shecan DNS servers for potentially faster and more secure internet browsing.
- **Reset DNS:** Revert to your original DNS settings with one click.
- **System Tray Integration:** Runs quietly in the system tray, allowing for easy access without cluttering your desktop.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Windows operating system with .NET Framework
- Administrator privileges for changing network settings

## Installation

1. Clone the repository to your local machine using `git clone`.
2. Navigate to the cloned directory.
3. Compile the application using a Golang compiler or directly run the binary if provided.

## Usage

To use PlaNet, follow these steps:

1. Run the application.
2. Right-click on the system tray icon.
3. Choose "Set" to change your DNS to Shecan or "Reset" to revert to the original settings.

Upon successful DNS change, the tooltip will display "Connected to: ...". If resetting, it will show "Not connected".

## Contributing

Contributions to the project are welcome. To contribute:

1. Fork the project.
2. Create a new branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

If you have any questions or suggestions, please reach out through GitHub issues.

---

**Note:** This tool changes system DNS settings. Use at your own risk. Always ensure you have backups of your original DNS configurations before making changes.