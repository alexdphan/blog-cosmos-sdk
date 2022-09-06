package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"

	// Keep track of the index of posts
	// To keep a list of posts in what is essentially a key-value store, you need to keep track of the index of the posts you insert. Since both post values and post count (index) values are kept in the store, you can use different prefixes: Post-value- and Post-count-.
	PostKey      = "Post-value-"
	PostCountKey = "Post-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ------------------------------------------------------------
// Your blog is now updated to take these actions when a Post message is sent to the AppendPost function:

// Get the number of posts in the store (count)
// Add a post by using the count as an ID
// Increment the count
// Return the count
