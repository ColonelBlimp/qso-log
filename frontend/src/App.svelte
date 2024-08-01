<script>
    import Panel from "./components/Panel.svelte";
    let name = 'Marc';
    let src = 'assets/img/kaningna.jpg';
    let gothtmlString = `This has <strong>html</strong> included!`;
    let count = 0;
    const increment = () => {
        count++;
    }
</script>

<div>Hello {name.toUpperCase()} test!</div>
<div><img src="{src}" alt="{name} lives here..." class="w-48 h-48"></div>
<Panel />
<p>{@html gothtmlString}</p>

<button on:click={increment}>
    Clicked {count}
    {count === 1 ? 'time' : 'times'}
</button>
