import asyncio
import logging
from aioquic.asyncio.client import connect
from aioquic.quic.configuration import QuicConfiguration

logging.basicConfig(level=logging.DEBUG)

async def main():
    config = QuicConfiguration(is_client=True)
    config.load_verify_locations("server.crt")

    async with connect("35.220.136.70", 31817, configuration=config) as client:
        h3_conn = client.http
        stream_id = h3_conn.get_next_available_stream_id()
        h3_conn.send_headers(
            stream_id=stream_id,
            headers=[
                (b":method", b"POST"),
                (b":scheme", b"https"),
                (b":authority", b"35.220.136.70"),
                (b":path", b"/"),
                (b"content-length", b"11"),
            ],
        )
        h3_conn.send_data(stream_id, b"I want flag", end_stream=True)

        while True:
            event = await client.wait_for_event()
            if hasattr(event, "data"):
                print(event.data.decode())
                break

asyncio.run(main())
