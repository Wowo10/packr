<script lang="ts">
    import ConfirmButton from "./common/ConfirmButton.svelte";

    let amountInput = 0;
    let solution: any = {}; //TODO: Add typing

    const getSolution = async () => {
        try {
            const response = await fetch(
                "/api/solution?amount=" + amountInput,
                {
                    headers: {
                        "X-Api-Key": import.meta.env.VITE_API_KEY,
                    },
                },
            );
            if (response.ok) {
                const data = await response.json();
                solution = data.solution;
            } else {
                console.error("Failed to fetch solution:", response.status);
            }
        } catch (error) {
            console.error("Failed to fetch solution:", error);
        }
    };
</script>

<div>
    <h2>Solve</h2>
    <input type="number" bind:value={amountInput} min="1" />
    <button on:click={() => getSolution()}>Solve</button>

    <div>
        {#each Object.entries(solution) as [pack, count]}
            <div style="margin-top: 5px;">
                <span>{count}</span>
                <span>x</span>
                <span>{pack}</span>
            </div>
        {/each}
    </div>
</div>

<style>
</style>
