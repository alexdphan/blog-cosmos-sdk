// define the AppendPost keeper method

// Create the x/blog/keeper/post.go file and start thinking about the logic of the function and what you want to call the prefixes. The file will be empty for now.

// To implement AppendPost you must first understand how the key store works. You can think of a store as a key-value database where keys are lexicographically ordered. You can loop through keys and use Get and Set to retrieve and set values based on keys. To distinguish between different types of data that a module can keep in its store, you can use prefixes like product- or post-.

// To keep a list of posts in what is essentially a key-value store, you need to keep track of the index of the posts you insert. Since both post values and post count (index) values are kept in the store, you can use different prefixes: Post-value- and Post-count-.

// In the x/blog/keeper/post.go file, draft the AppendPost function. You can add these comments to help you visualize what you do next:

package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// Get the current number of posts in the store
	count := k.GetPostCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	post.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetPostCount(ctx, count+1)
	return count
}

// Implementing the GetPostCount function
func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Now that GetPostCount returns the correct number of posts in the store, implement SetPostCount:
func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	// Convert the PostCountKey (which is "Post-count-") to bytes
	byteKey := []byte(types.PostCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Post-count- (byteKey) to count
	store.Set(byteKey, bz)
}

// By following these steps, you have implemented all of the code required to create new posts and store them on-chain. Now, when a transaction that contains a message of type MsgCreatePost is broadcast, the message is routed to your blog module.

// - k.CreatePost calls AppendPost
// - AppendPost gets the number of posts from the store, adds a post using the count as an ID, increments the count, and returns the ID

// Now that you have added the functionality to create posts and broadcast them to our chain, you can add querying.
