<script lang="ts">
    import load from "../assets/load.gif";

    let amountInput = 0;
    let solution: any = {}; //TODO: Add typing
    let loading = false;
    const baseurl = import.meta.env.VITE_API_BASE_URL;

    const getSolution = async () => {
        try {
            loading = true;
            const response = await fetch(
                baseurl + "/api/solution?amount=" + amountInput,
                {
                    headers: {
                        "X-Api-Key": import.meta.env.VITE_API_KEY,
                    },
                },
            );
            if (response.ok) {
                const data = await response.json();
                solution = data.solution;
                loading = false;
            } else {
                console.error("Failed to fetch solution:", response.status);
            }
        } catch (error) {
            console.error("Failed to fetch solution:", error);
            loading = false;
        }
    };
</script>

<div>
    <h2>Solve</h2>
    <input type="number" bind:value={amountInput} min="1" />
    <button on:click={() => getSolution()}>Solve</button>
    {#if loading}
        <img src={load} alt="loading" width="20px" height="20px" />
    {/if}

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
