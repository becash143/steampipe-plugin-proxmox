connection "proxmox" {
  plugin = "becash143/proxmox"

  # The base URL of your Proxmox VE host, without any API path suffix.
  # This typically points at port 8006. Do not include /api2/json —
  # the plugin appends it automatically to every request.
  # Can also be set with the PROXMOX_ENDPOINT environment variable.
  endpoint = "https://pve.example.com:8006"

  # API token ID, in the format "user@realm!tokenid".
  # Create one in the Proxmox UI under Datacenter > Permissions > API Tokens.
  # Can also be set with the PROXMOX_API_TOKEN environment variable.
  api_token = "root@pam!steampipe"

  # The secret value generated when the API token above was created.
  # Proxmox only shows this once, so store it securely (e.g. a secrets manager).
  # Can also be set with the PROXMOX_API_SECRET environment variable.
  api_secret = "1234abcd-5678-90ef-abcd-1234567890ab"

  # Set to true to skip TLS certificate verification, e.g. when connecting
  # to a Proxmox node using a self-signed certificate. Defaults to false.
  # Only disable verification in trusted, internal network environments.
  insecure = false
}
