package commandengine

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"example.com/command/http"
	"example.com/command/posts"
	"example.com/command/writer"
	"github.com/spf13/cobra"
)

type CommandEngine struct {
	rootCmd *cobra.Command
}

func NewCommandEngine() *CommandEngine {
	rootCmd := &cobra.Command{
		Use:   "help",
		Short: "my demo command line application",
		Long:  "commands for my sample command line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("executed for help")
		},
	}

	defer func() {
		// if the init results in panic then capture it
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	return &CommandEngine{
		rootCmd: rootCmd,
	}
}

func (c *CommandEngine) Run() error {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "..")
	commands := []*cobra.Command{
		{
			Use:   "execute",
			Short: "start the http server",
			Long:  "command to start the server to execute something",
			Run: func(cmd *cobra.Command, args []string) {
				p := &posts.Posts{
					PostGetter: http.NewHTTP("https://jsonplaceholder.typicode.com"),
					PostWriter: writer.NewWriter(),
				}
				postID, _ := strconv.ParseInt(args[0], 10, 32)
				post, _ := p.GetPostForAnId(int(postID))
				err := p.PostWriter.WriteToTxt(fmt.Sprintf("metadata/data_id_%d", postID), root, post.ToString())
				if err != nil {
					fmt.Printf("failed to write to file %s", err.Error())
				}
			},
		},
	}

	for _, command := range commands {
		c.rootCmd.AddCommand(command)
	}
	err := c.rootCmd.Execute()
	return err
}
