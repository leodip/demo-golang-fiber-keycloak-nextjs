import { NextResponse } from "next/server";
import { getAccessToken } from "../../../utils/sessionTokenAccessor";
import { getServerSession } from "next-auth";
import { authOptions } from "../auth/[...nextauth]/route";

export async function POST(req) {
  const session = await getServerSession(authOptions);

  if (session) {
    const url = `${process.env.DEMO_BACKEND_URL}/api/v1/products`;

    const postBody = await req.json();
    let accessToken = await getAccessToken();

    const resp = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + accessToken,
      },
      method: "POST",
      body: JSON.stringify(postBody),
    });

    if (resp.ok) {
      const data = await resp.json();
      return NextResponse.json({ data }, { status: resp.status });
    }
    
    return NextResponse.json(
      { error: await resp.text() },
      { status: resp.status }
    );
  }
  return NextResponse.json({ error: "Unauthorized" }, { status: res.status });
}