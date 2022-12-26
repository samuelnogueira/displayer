Janky Discord bot to notify players entering/leaving multiplayer servers.

Currently offers support for:
* Minecraft
* Farming Simulator 22 Dedicated Server

## Environment Variables
| Name                   | Required | Description                                                                     | Example                                                                       |
|------------------------|----------|---------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| `DISCORD_WEBHOOK_URL`  | Yup      | Discord webhook URL. Will be used to send messages to a channel.                | `https://discord.com/api/webhooks/12345678/hlTyVj5NOVUV_RGuNhV51S02cFzBs5xT6` |
| `MINECRAFT_SERVER_URL` | Nope      | Minecraft server host. Will be pinged every minute to poll player list changes. | `//minecraft.example.com.:25565`                                              |
| `FS22_SERVER_STATS_URL` | No      | Farming Simulator 22 stats XML URL. Will be pinged every minute to poll player list changes. | `https://fs22.example.com/feed/dedicated-server-stats.xml`                                              |
