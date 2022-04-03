# discraft
Janky Discord bot to notify players entering/leaving a Minecraft server.

## Environment Variables
| Name                   | Required | Description                                                                     | Example                                                                       |
|------------------------|----------|---------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| `DISCORD_WEBHOOK_URL`  | Yup      | Discord webhook URL. Will be used to send messages to a channel.                | `https://discord.com/api/webhooks/12345678/hlTyVj5NOVUV_RGuNhV51S02cFzBs5xT6` |
| `MINECRAFT_SERVER_URL` | Yep      | Minecraft server host. Will be pinged every minute to poll player list changes. | `//minecraft.example.com.:25565`                                              |
