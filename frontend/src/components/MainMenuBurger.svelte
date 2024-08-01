<script>
    import { mainMenuStore } from "../lib/stores.js";

    let isMainMenuOpen = false;
    mainMenuStore.subscribe((value) => {
        isMainMenuOpen = value;
    });

    const menuToggle = () => {
        if (isMainMenuOpen) {
            mainMenuStore.update((b) => false);
        } else {
            mainMenuStore.update((b) => true);
        }
    }
</script>
<button on:click={menuToggle} type="button" class="relative inline-flex items-center justify-center rounded-md bg-indigo-600 p-2 text-indigo-200 hover:bg-indigo-500 hover:bg-opacity-75 hover:text-white focus:outline-none focus:ring-2 focus:ring-white/50 focus:ring-offset-2 focus:ring-offset-indigo-600" aria-controls="mainMenu">
    {#if !isMainMenuOpen}
    <svg class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
    </svg>
    {:else}
    <svg class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
    </svg>
    {/if}
</button>
