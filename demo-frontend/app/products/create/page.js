"use client";

import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import React, { useState, useEffect } from "react";

export default function CreateProduct() {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (
      status == "unauthenticated" ||
      (status == "authenticated" && !session.roles?.includes("admin"))
    ) {
      router.push("/unauthorized");
      router.refresh();
    }
  }, [session, status, router]);

  const productNameRef = React.useRef();
  const priceRef = React.useRef();

  const [errorMsg, setErrorMsg] = useState("");

  if (status == "loading") {
    return (
      <main>
        <h1 className="text-4xl text-center">Create product</h1>
        <div className="text-center text-2xl">Loading...</div>
      </main>
    );
  }

  if (session && session.roles?.includes("admin")) {
    const handleSubmit = async (event) => {
      event.preventDefault();

      const postBody = {
        Name: productNameRef.current.value,
        Price: parseFloat(priceRef.current.value),
      };

      try {
        const resp = await fetch("/api/products", {
          method: "POST",
          headers: {
            headers: {
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify(postBody),
        });

        if (resp.ok) {
          router.push("/products");
          router.refresh();
        } else {
          var json = await resp.json();
          setErrorMsg("Unable to call the API: " + json.error);
        }
      } catch (err) {
        setErrorMsg("Unable to call the API: " + err);
      }
    };

    return (
      <main>
        <h1 className="text-4xl text-center">Create product</h1>

        <form onSubmit={handleSubmit} className="mt-6">
          <div className="w-1/2">
            <label htmlFor="productName" className="text-2xl">Product name:</label>
            <input autoFocus type="text" id="productName" 
                className="w-full p-1 text-black bg-gray-200 text-lg" ref={productNameRef} required />
          </div>
          <div className="w-1/2 mt-2">
            <label htmlFor="price" className="text-2xl">
              Price:
            </label>
            <input type="number" step="0.01" id="price" className="w-full p-1 text-black bg-gray-200 text-lg" ref={priceRef} />
          </div>
          <div className="text-center text-2xl text-red-600">{errorMsg}</div>
          <button type="submit" className="mt-3 bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-50">
            Create
          </button>
        </form>
      </main>
    );
  }
}