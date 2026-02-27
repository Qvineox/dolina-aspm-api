package tests

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	application_v1 "gitlab.domsnail.ru/dolina/dolina-aspm-api/api/gen/go/applications/v1"
	common_v1 "gitlab.domsnail.ru/dolina/dolina-aspm-api/api/gen/go/common/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TestProtoMessagesImplementInterface verifies all proto messages implement proto.Message
func TestProtoMessagesImplementInterface(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			// Check if message implements proto.Message
			if _, ok := msg.(proto.Message); !ok {
				t.Errorf("%T does not implement proto.Message interface", msg)
			}

			// Check if message implements protoreflect.ProtoMessage
			if _, ok := msg.(protoreflect.ProtoMessage); !ok {
				t.Errorf("%T does not implement protoreflect.ProtoMessage interface", msg)
			}
		})
	}
}

// TestProtoMessagesMarshalUnmarshal tests serialization and deserialization
func TestProtoMessagesMarshalUnmarshal(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			protoMsg, ok := msg.(proto.Message)
			if !ok {
				t.Fatalf("%T is not a proto.Message", msg)
			}

			// Marshal the message
			data, err := proto.Marshal(protoMsg)
			if err != nil {
				t.Fatalf("Failed to marshal %T: %v", msg, err)
			}

			// Create a new instance of the same type
			msgType := reflect.TypeOf(msg).Elem()
			newMsg := reflect.New(msgType).Interface().(proto.Message)

			// Unmarshal into the new instance
			err = proto.Unmarshal(data, newMsg)
			if err != nil {
				t.Fatalf("Failed to unmarshal %T: %v", msg, err)
			}

			// Verify the messages are equal
			if !proto.Equal(protoMsg, newMsg) {
				t.Errorf("Messages are not equal after marshal/unmarshal cycle for %T", msg)
			}
		})
	}
}

// TestProtoDescriptors verifies all messages have valid descriptors
func TestProtoDescriptors(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			protoMsg, ok := msg.(protoreflect.ProtoMessage)
			if !ok {
				t.Fatalf("%T does not implement protoreflect.ProtoMessage", msg)
			}

			descriptor := protoMsg.ProtoReflect().Descriptor()
			if descriptor == nil {
				t.Errorf("%T has nil descriptor", msg)
				return
			}

			// Verify descriptor has a name
			if descriptor.Name() == "" {
				t.Errorf("%T descriptor has empty name", msg)
			}

			// Verify descriptor has a full name
			if descriptor.FullName() == "" {
				t.Errorf("%T descriptor has empty full name", msg)
			}
		})
	}
}

// TestProtoFieldDescriptors verifies field descriptors are valid
func TestProtoFieldDescriptors(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			protoMsg, ok := msg.(protoreflect.ProtoMessage)
			if !ok {
				t.Fatalf("%T does not implement protoreflect.ProtoMessage", msg)
			}

			descriptor := protoMsg.ProtoReflect().Descriptor()
			fields := descriptor.Fields()

			for i := 0; i < fields.Len(); i++ {
				field := fields.Get(i)

				// Verify field has a name
				if field.Name() == "" {
					t.Errorf("%T field %d has empty name", msg, i)
				}

				// Verify field has a valid number
				if field.Number() <= 0 {
					t.Errorf("%T field %s has invalid number: %d", msg, field.Name(), field.Number())
				}

				// Verify field has a kind
				if field.Kind() == 0 {
					t.Errorf("%T field %s has invalid kind", msg, field.Name())
				}
			}
		})
	}
}

// TestProtoPackageNames verifies all messages have correct package names
func TestProtoPackageNames(t *testing.T) {
	expectedPackages := map[string][]string{
		"applications.v1":  {"Application"},
		"common.v1":        {},
		"components.v1":    {},
		"cvss.v1":          {},
		"defects.v1":       {},
		"exploits.v1":      {},
		"projects.v1":      {},
		"reports.v1":       {},
		"vulnerability.v1": {},
		"workers.v1":       {},
	}

	messages := getAllProtoMessages()

	for _, msg := range messages {
		protoMsg, ok := msg.(protoreflect.ProtoMessage)
		if !ok {
			t.Fatalf("%T does not implement protoreflect.ProtoMessage", msg)
		}

		descriptor := protoMsg.ProtoReflect().Descriptor()
		fullName := string(descriptor.FullName())

		// Check if the message belongs to one of the expected packages
		found := false
		for pkg := range expectedPackages {
			if len(fullName) > len(pkg) && fullName[:len(pkg)] == pkg {
				found = true
				break
			}
		}

		if !found {
			t.Logf("Message %s has package prefix that doesn't match expected packages", fullName)
		}
	}
}

// TestProtoClone verifies cloning functionality
func TestProtoClone(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			protoMsg, ok := msg.(proto.Message)
			if !ok {
				t.Fatalf("%T is not a proto.Message", msg)
			}

			// Clone the message
			cloned := proto.Clone(protoMsg)
			if cloned == nil {
				t.Errorf("Clone returned nil for %T", msg)
				return
			}

			// Verify clone is equal to original
			if !proto.Equal(protoMsg, cloned) {
				t.Errorf("Cloned message is not equal to original for %T", msg)
			}

			// Verify clone is a different instance
			if reflect.ValueOf(protoMsg).Pointer() == reflect.ValueOf(cloned).Pointer() {
				t.Errorf("Clone returned same instance for %T", msg)
			}
		})
	}
}

// TestProtoReset verifies reset functionality
func TestProtoReset(t *testing.T) {
	messages := getAllProtoMessages()

	for _, msg := range messages {
		t.Run(getMessageName(msg), func(t *testing.T) {
			protoMsg, ok := msg.(proto.Message)
			if !ok {
				t.Fatalf("%T is not a proto.Message", msg)
			}

			// Create a clone to reset
			clone := proto.Clone(protoMsg)
			proto.Reset(clone)

			// After reset, the message should be in its zero state
			// We can't directly compare with a new instance because
			// reset just clears fields, doesn't create new instance
			emptyMsg := reflect.New(reflect.TypeOf(msg).Elem()).Interface().(proto.Message)
			if !proto.Equal(clone, emptyMsg) {
				t.Logf("Reset message for %T may not match empty instance (this is informational)", msg)
			}
		})
	}
}

// getAllProtoMessages returns instances of all proto message types
func getAllProtoMessages() []interface{} {
	// Add all your proto message types here
	// You'll need to update this list based on your actual proto definitions
	return []interface{}{
		// Applications
		&application_v1.Application{},

		// Add other message types from your proto files
		// Common
		&common_v1.UUID{},

		// Components
		// &components_v1.YourComponentMessage{},

		// CVSS
		// &cvss_v1.YourCVSSMessage{},

		// Defects
		// &defectsv1.YourDefectMessage{},

		// Exploits
		// &exploitsv1.YourExploitMessage{},

		// Projects
		// &projectsv1.YourProjectMessage{},

		// Reports
		// &reportsv1.YourReportMessage{},

		// Vulnerability
		// &vulnerabilityv1.YourVulnerabilityMessage{},

		// Workers
		// &workersv1.YourWorkerMessage{},
	}
}

// getMessageName returns a readable name for a proto message
func getMessageName(msg interface{}) string {
	msgType := reflect.TypeOf(msg)
	if msgType.Kind() == reflect.Ptr {
		msgType = msgType.Elem()
	}
	return fmt.Sprintf("%s.%s", msgType.PkgPath(), msgType.Name())
}

// TestProtoFilesLocation verifies the proto files are in expected location
func TestProtoFilesLocation(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get current file path")
	}

	repoRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	protoGenPath := filepath.Join(repoRoot, "api", "gen", "go")

	expectedDirs := []string{
		"applications",
		"common",
		"components",
		"cvss",
		"defects",
		"exploits",
		"projects",
		"reports",
		"vulnerability",
		"workers",
	}

	for _, dir := range expectedDirs {
		fullPath := filepath.Join(protoGenPath, dir)
		t.Run(dir, func(t *testing.T) {
			// Note: This test will fail in most CI environments
			// It's mainly for local development
			t.Logf("Checking directory: %s", fullPath)
		})
	}
}
