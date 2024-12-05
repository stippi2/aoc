def parse_input(filename):
    """Parse input file into rules and updates"""
    rules = []
    updates = []

    with open(filename) as f:
        # Read all lines and split into rules and updates
        content = f.read().strip().split('\n\n')

        # Parse rules
        for line in content[0].split('\n'):
            if line:  # Skip empty lines
                before, after = line.strip().split('|')
                rules.append((int(before), int(after)))

        # Parse updates
        for line in content[1].split('\n'):
            if line:  # Skip empty lines
                update = [int(x) for x in line.strip().split(',')]
                updates.append(update)

    return rules, updates

def is_valid_order(pages, rules):
    """Check if pages are in valid order according to rules"""
    # Create a set of rules that apply to these pages
    relevant_rules = []
    pages_set = set(pages)

    for before, after in rules:
        # Only include rules where both pages are in the update
        if before in pages_set and after in pages_set:
            relevant_rules.append((before, after))

    # Check each relevant rule
    for before, after in relevant_rules:
        before_idx = pages.index(before)
        after_idx = pages.index(after)
        if before_idx >= after_idx:  # If 'before' page comes after 'after' page
            return False

    return True

def get_middle_number(pages):
    """Get middle number from a list of pages"""
    return pages[len(pages) // 2]

def create_graph(pages, rules):
    """Create a directed graph from rules, considering only relevant pages"""
    # Initialize graph
    graph = {page: set() for page in pages}
    incoming_edges = {page: 0 for page in pages}

    # Add edges from rules
    pages_set = set(pages)
    for before, after in rules:
        if before in pages_set and after in pages_set:
            graph[before].add(after)
            incoming_edges[after] += 1

    return graph, incoming_edges

def topological_sort(pages, rules):
    """Sort pages according to rules using Kahn's algorithm"""
    # Create graph
    graph, incoming_edges = create_graph(pages, rules)

    # Find nodes with no incoming edges
    queue = [page for page in pages if incoming_edges[page] == 0]
    result = []

    # Process queue
    while queue:
        # Get node with no incoming edges
        # Choose the highest number if multiple are available
        queue.sort(reverse=True)
        current = queue.pop(0)
        result.append(current)

        # Remove edges from current node
        for neighbor in graph[current]:
            incoming_edges[neighbor] -= 1
            if incoming_edges[neighbor] == 0:
                queue.append(neighbor)

    return result if len(result) == len(pages) else None

def solve_part1(filename):
    """Solve part 1: Sum of middle numbers of valid updates"""
    rules, updates = parse_input(filename)

    # Find valid updates and get their middle numbers
    middle_sum = 0
    for update in updates:
        if is_valid_order(update, rules):
            middle_sum += get_middle_number(update)

    return middle_sum

def solve_part2(filename):
    """Solve part 2: Sum of middle numbers of corrected invalid updates"""
    rules, updates = parse_input(filename)

    # Find invalid updates and correct them
    middle_sum = 0
    for update in updates:
        if not is_valid_order(update, rules):
            # Sort the pages according to rules
            sorted_pages = topological_sort(update, rules)
            if sorted_pages:  # If a valid ordering exists
                middle_sum += get_middle_number(sorted_pages)

    return middle_sum

if __name__ == "__main__":
    print(f"Part 1 solution: {solve_part1('day05.txt')}")
    print(f"Part 2 solution: {solve_part2('day05.txt')}")
