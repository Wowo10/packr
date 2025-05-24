<script lang="ts">
    import { onMount } from "svelte";
    import ConfirmButton from "./common/ConfirmButton.svelte";

    let packs: number[] = [];
    let addPackInput = 0;

    const getPacks = async () => {
        try {
            const response = await fetch("/api/packs", {
                headers: {
                    "X-Api-Key": import.meta.env.VITE_API_KEY,
                },
            });
            if (response.ok) {
                const data = await response.json();
                packs = data.packs;
            } else {
                console.error("Failed to fetch packs:", response.status);
            }
        } catch (error) {
            console.error("Failed to fetch packs:", error);
        }
    };

    onMount(getPacks);

    const deletePack = async (pack: number) => {
        try {
            const res = await fetch("/api/packs?pack=" + pack, {
                method: "DELETE",
                headers: {
                    "X-Api-Key": import.meta.env.VITE_API_KEY,
                },
            });

            if (res.ok) {
                getPacks();
            } else {
                console.error("Failed to remove pack:", await res.text());
            }
        } catch (e) {
            console.error("Failed to remove pack:", e);
        }
    };

    const addPack = async (pack: number) => {
        try {
            const res = await fetch("/api/packs?pack=" + pack, {
                method: "POST",
                headers: {
                    "X-Api-Key": import.meta.env.VITE_API_KEY,
                },
            });

            if (res.ok) {
                getPacks();
            } else {
                console.error("Failed to remove pack:", await res.text());
            }
        } catch (e) {
            console.error("Failed to remove pack:", e);
        }
    };

    const handleKeyDown = (e: KeyboardEvent) => {
        if (
            e.key === "Enter" &&
            document.activeElement === (e.target as HTMLInputElement)
        ) {
            addPack(addPackInput);
        }
    };

    document.addEventListener("keydown", handleKeyDown);
</script>

<div>
    <h2>Packs</h2>
    {#each packs as pack}
        <div style="margin-top: 5px;">
            <span class="packContainer">{pack}</span>
            <ConfirmButton
                onConfirm={() => deletePack(pack)}
                label="Remove"
                confirmLabel="Sure?"
            />
        </div>
    {/each}
    <div>
        <input type="number" bind:value={addPackInput} min="1" />
        <ConfirmButton
            onConfirm={() => addPack(addPackInput)}
            label="Add"
            confirmLabel="Sure?"
        />
    </div>
</div>

<style>
    .packContainer {
        border: 2px black solid;
        border-radius: 5px;
        padding: 5px 5px;
    }
</style>
