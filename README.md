# discraft
Janky Discord bot to notify players entering/leaving a Minecraft server.

## Environment Variables
| Name | Required | Description | Example |
| ---- | -------- | ----------- | --------|
| `DISCORD_WEBHOOK_URL` | Yup | Discord Webhook URL (will be used to send messages to a channel) | `https://discord.com/api/webhooks/1234567891236/hlTyFb2OLU8HVg4mhZ4sLrr3Vj5NOVUV_RGuNhV51S02cyZXbVnWJiAIVFzBs5xT6` |
| `MINECRAFT_HOST` | Yep | Minecraft server host (without port) | `minecraft.example.com` |
| `MINECRAFT_PORT` | Yes | Minecraft server port | `25565` |
